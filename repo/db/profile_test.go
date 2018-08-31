package db

import (
	"database/sql"
	"github.com/textileio/textile-go/repo"
	"sync"
	"testing"
	"time"
)

var pdb repo.ProfileStore

func init() {
	setupProfileDB()
}

func setupProfileDB() {
	conn, _ := sql.Open("sqlite3", ":memory:")
	initDatabaseTables(conn, "")
	pdb = NewProfileStore(conn, new(sync.Mutex))
}

func TestProfileDB_GetTokensPreSignIn(t *testing.T) {
	tokens, err := pdb.GetTokens()
	if err != nil {
		t.Error(err)
		return
	}
	if tokens != nil {
		t.Error("tokens should be nil")
	}
}

func TestProfileDB_SignIn(t *testing.T) {
	err := pdb.SignIn("woohoo!", &repo.CafeTokens{Access: "access", Refresh: "refresh", Expiry: time.Now()})
	if err != nil {
		t.Error(err)
	}
}

func TestProfileDB_GetUsername(t *testing.T) {
	un, err := pdb.GetUsername()
	if err != nil {
		t.Error(err)
		return
	}
	if *un != "woohoo!" {
		t.Error("got bad username")
	}
}

func TestProfileDB_SetAvatarId(t *testing.T) {
	if err := pdb.SetAvatarId("/ipfs/Qm..."); err != nil {
		t.Error(err)
		return
	}
}

func TestProfileDB_GetAvatarId(t *testing.T) {
	av, err := pdb.GetAvatarId()
	if err != nil {
		t.Error(err)
		return
	}
	if *av != "/ipfs/Qm..." {
		t.Error("got bad avatar id")
	}
}

func TestProfileDB_GetTokens(t *testing.T) {
	tokens, err := pdb.GetTokens()
	if err != nil {
		t.Error(err)
		return
	}
	if tokens.Access != "access" {
		t.Error("got bad access token")
		return
	}
	if tokens.Refresh != "refresh" {
		t.Error("got bad refresh token")
		return
	}
}

func TestProfileDB_UpdateTokens(t *testing.T) {
	err := pdb.UpdateTokens(&repo.CafeTokens{Access: "access", Refresh: "refresh", Expiry: time.Now()})
	if err != nil {
		t.Error(err)
	}
}

func TestProfileDB_SignOut(t *testing.T) {
	if err := pdb.SignOut(); err != nil {
		t.Error(err)
		return
	}
	tokens, err := pdb.GetTokens()
	if err != nil {
		t.Error(err)
	}
	if tokens != nil {
		t.Error("signed out but tokens still present")
	}
}
