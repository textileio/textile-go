package core

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	cid "gx/ipfs/QmTbxNB1NwDesLmKTscr4udL2tVP7MaxvXnD1D9yX7g3PN/go-cid"
	"gx/ipfs/QmYVXrKrKHDC9FobgmcmshCDyWwdrfwfanNQN4oxJ9Fk3h/go-libp2p-peer"
	uio "gx/ipfs/QmcYUTQ7tBZeH1CLsZM2S3xhMEZdvUgXvbjhpMsLDpk3oJ/go-unixfs/io"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/golang/protobuf/proto"
	"github.com/textileio/go-textile/ipfs"
	"github.com/textileio/go-textile/pb"
)

// pin take raw data or a tarball and pins it to the local ipfs node.
// request must be authenticated with a token
func (c *cafeApi) pin(g *gin.Context) {
	// handle based on content type
	var id cid.Cid
	cType := g.Request.Header.Get("Content-Type")
	switch cType {
	case "application/gzip":
		dirb := uio.NewDirectory(c.node.Ipfs().DAG)

		gr, err := gzip.NewReader(g.Request.Body)
		if err != nil {
			log.Errorf("error creating gzip reader %s", err)
			c.abort(g, http.StatusBadRequest, err)
			return
		}
		tr := tar.NewReader(gr)

		for {
			header, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Errorf("error getting tar next %s", err)
				c.abort(g, http.StatusInternalServerError, err)
				return
			}

			switch header.Typeflag {
			case tar.TypeDir:
				log.Error("got nested directory, aborting")
				c.abort(g, http.StatusBadRequest, fmt.Errorf("directories are not supported"))
				return
			case tar.TypeReg:
				if _, err := ipfs.AddDataToDirectory(c.node.Ipfs(), dirb, header.Name, tr); err != nil {
					log.Errorf("error adding file to dir %s", err)
					c.abort(g, http.StatusInternalServerError, err)
					return
				}
			default:
				continue
			}
		}

		// pin the directory
		dir, err := dirb.GetNode()
		if err != nil {
			log.Errorf("error creating dir node %s", err)
			c.abort(g, http.StatusInternalServerError, err)
			return
		}
		if err := ipfs.PinNode(c.node.Ipfs(), dir, true); err != nil {
			log.Errorf("error pinning dir node %s", err)
			c.abort(g, http.StatusInternalServerError, err)
			return
		}
		id = dir.Cid()

	case "application/octet-stream":
		idp, err := ipfs.AddData(c.node.Ipfs(), g.Request.Body, true)
		if err != nil {
			log.Errorf("error pinning raw body %s", err)
			c.abort(g, http.StatusInternalServerError, err)
			return
		}
		id = *idp
	default:
		log.Errorf("got bad content type %s", cType)
		c.abort(g, http.StatusBadRequest, fmt.Errorf("invalid content-type"))
		return
	}
	hash := id.Hash().B58String()

	log.Debugf("pinned request with content type %s: %s", cType, hash)

	g.JSON(http.StatusCreated, gin.H{"id": hash})
}

// service is an HTTP entry point for the cafe service
func (c *cafeApi) service(g *gin.Context) {
	body, err := ioutil.ReadAll(g.Request.Body)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	// parse body as a service envelope
	pmes := new(pb.Envelope)
	if err := proto.Unmarshal(body, pmes); err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	peerId := g.Request.Header.Get("X-Textile-Peer")
	if peerId == "" {
		g.String(http.StatusBadRequest, "missing peer ID")
		return
	}
	mPeer, err := peer.IDB58Decode(peerId)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := c.node.cafe.service.VerifyEnvelope(pmes, mPeer); err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	// handle the message as normal
	log.Debugf("received %s from %s", pmes.Message.Type.String(), mPeer.Pretty())
	rpmes, err := c.node.cafe.Handle(mPeer, pmes)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	if rpmes != nil {
		res, err := proto.Marshal(rpmes)
		if err != nil {
			g.String(http.StatusBadRequest, err.Error())
			return
		}

		g.Render(http.StatusOK, render.Data{Data: res})
		return
	}

	// handle the message as a JSON stream
	rpmesCh, errCh, cancel := c.node.cafe.HandleStream(mPeer, pmes)
	g.Stream(func(w io.Writer) bool {
		select {
		case <-g.Request.Context().Done():
			close(cancel)

		case err := <-errCh:
			g.String(http.StatusBadRequest, err.Error())
			return false

		case rpmes, ok := <-rpmesCh:
			if !ok {
				g.Status(http.StatusOK)
				return false
			}
			log.Debugf("responding with %s to %s", rpmes.Message.Type.String(), mPeer.Pretty())

			g.JSON(http.StatusOK, rpmes)
			g.Writer.Write([]byte("\n"))
		}
		return true
	})
}
