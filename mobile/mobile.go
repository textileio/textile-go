package mobile

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/op/go-logging"
	"github.com/textileio/textile-go/central/models"
	tcore "github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/crypto"
	"github.com/textileio/textile-go/repo"
	"github.com/textileio/textile-go/wallet"
	"github.com/textileio/textile-go/wallet/thread"
	"github.com/textileio/textile-go/wallet/util"
	libp2pc "gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"
	"time"
)

var log = logging.MustGetLogger("mobile")

// Message is a generic go -> bridge message structure
type Event struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
}

// newEvent transforms an event name and structured data in Event
func newEvent(name string, payload map[string]interface{}) *Event {
	event := &Event{Name: name}
	jsn, _ := toJSON(payload)
	event.Payload = jsn
	return event
}

// Messenger is used to inform the bridge layer of new data waiting to be queried
type Messenger interface {
	Notify(event *Event)
}

// NodeConfig is used to configure the mobile node
// NOTE: logLevel is one of: CRITICAL ERROR WARNING NOTICE INFO DEBUG
type NodeConfig struct {
	RepoPath      string
	CentralApiURL string
	LogLevel      string
	LogFiles      bool
}

// Mobile is the name of the framework (must match package name)
type Mobile struct {
	RepoPath  string
	Mnemonic  string
	Online    <-chan struct{} // not readable from bridges
	messenger Messenger
}

// Thread is a simple meta data wrapper around a Thread
type Thread struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Peers int    `json:"peers"`
}

// Threads is a wrapper around a list of Threads
type Threads struct {
	Items []Thread `json:"items"`
}

// Device is a simple meta data wrapper around a Device
type Device struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Devices is a wrapper around a list of Devices
type Devices struct {
	Items []Device `json:"items"`
}

// Photo is a simple meta data wrapper around a photo block
type Photo struct {
	Id       string    `json:"id"`
	Date     time.Time `json:"date"`
	AuthorId string    `json:"author_id"`
	Caption  string    `json:"caption"`
}

// Photos is a wrapper around a list of photos
type Photos struct {
	Items []Photo `json:"items"`
}

// ExternalInvite is a wrapper around an invite id and key
type ExternalInvite struct {
	Id      string `json:"id"`
	Key     string `json:"key"`
	Inviter string `json:"inviter"`
}

// Create a gomobile compatible wrapper around TextileNode
func NewNode(config *NodeConfig, messenger Messenger) (*Mobile, error) {
	ll, err := logging.LogLevel(config.LogLevel)
	if err != nil {
		ll = logging.INFO
	}
	cconfig := tcore.NodeConfig{
		LogLevel: ll,
		LogFiles: config.LogFiles,
		WalletConfig: wallet.Config{
			RepoPath:   config.RepoPath,
			CentralAPI: config.CentralApiURL,
			IsMobile:   true,
		},
	}
	node, mnemonic, err := tcore.NewNode(cconfig)
	if err != nil {
		return nil, err
	}
	tcore.Node = node

	return &Mobile{RepoPath: config.RepoPath, Mnemonic: mnemonic, messenger: messenger}, nil
}

// Start the mobile node
func (m *Mobile) Start() error {
	online, err := tcore.Node.StartWallet()
	if err != nil {
		if err == wallet.ErrStarted {
			return nil
		}
		return err
	}
	m.Online = online

	go func() {
		<-online
		// subscribe to thread updates
		for _, thrd := range tcore.Node.Wallet.Threads() {
			go func(t *thread.Thread) {
				m.subscribe(t)
			}(thrd)
		}

		// notify UI we're ready
		m.messenger.Notify(newEvent("onOnline", map[string]interface{}{}))

		// publish
		//tcore.Node.Wallet.PublishThreads()
	}()

	return nil
}

// Stop the mobile node
func (m *Mobile) Stop() error {
	if err := tcore.Node.StopWallet(); err != nil && err != wallet.ErrStopped {
		return err
	}
	return nil
}

// SignUpWithEmail creates an email based registration and calls core signup
func (m *Mobile) SignUpWithEmail(username string, password string, email string, referral string) error {
	// build registration
	reg := &models.Registration{
		Username: username,
		Password: password,
		Identity: &models.Identity{
			Type:  models.EmailAddress,
			Value: email,
		},
		Referral: referral,
	}
	return tcore.Node.Wallet.SignUp(reg)
}

// SignIn build credentials and calls core SignIn
func (m *Mobile) SignIn(username string, password string) error {
	// build creds
	creds := &models.Credentials{
		Username: username,
		Password: password,
	}
	return tcore.Node.Wallet.SignIn(creds)
}

// SignOut calls core SignOut
func (m *Mobile) SignOut() error {
	return tcore.Node.Wallet.SignOut()
}

// IsSignedIn calls core IsSignedIn
func (m *Mobile) IsSignedIn() bool {
	si, _ := tcore.Node.Wallet.IsSignedIn()
	return si
}

// GetId calls core GetId
func (m *Mobile) GetId() (string, error) {
	return tcore.Node.Wallet.GetId()
}

// GetUsername calls core GetUsername
func (m *Mobile) GetUsername() (string, error) {
	return tcore.Node.Wallet.GetUsername()
}

// GetAccessToken calls core GetAccessToken
func (m *Mobile) GetAccessToken() (string, error) {
	return tcore.Node.Wallet.GetAccessToken()
}

// Threads lists all threads
func (m *Mobile) Threads() (string, error) {
	threads := Threads{Items: make([]Thread, 0)}
	for _, thrd := range tcore.Node.Wallet.Threads() {
		peers := thrd.Peers()
		item := Thread{Id: thrd.Id, Name: thrd.Name, Peers: len(peers)}
		threads.Items = append(threads.Items, item)
	}
	return toJSON(threads)
}

// AddThread adds a new thread with the given name
func (m *Mobile) AddThread(name string, mnemonic string) (string, error) {
	var mnem *string
	if mnemonic != "" {
		mnem = &mnemonic
	}
	thrd, _, err := tcore.Node.Wallet.AddThreadWithMnemonic(name, mnem)
	if err != nil {
		return "", err
	}

	// subscribe to updates
	go m.subscribe(thrd)

	// build json
	peers := thrd.Peers()
	item := Thread{
		Id:    thrd.Id,
		Name:  thrd.Name,
		Peers: len(peers),
	}
	return toJSON(item)
}

// AddThreadInvite adds a new invite to a thread
func (m *Mobile) AddThreadInvite(threadId string, inviteePk string) (string, error) {
	_, thrd := tcore.Node.Wallet.GetThread(threadId)
	if thrd == nil {
		return "", errors.New(fmt.Sprintf("could not find thread: %s", threadId))
	}

	// decode pubkey
	ikb, err := libp2pc.ConfigDecodeKey(inviteePk)
	if err != nil {
		return "", err
	}
	ipk, err := libp2pc.UnmarshalPublicKey(ikb)
	if err != nil {
		return "", err
	}

	// add it
	addr, err := thrd.AddInvite(ipk)
	if err != nil {
		return "", err
	}

	return addr.B58String(), nil
}

// AddExternalThreadInvite generates a new external invite link to a thread
func (m *Mobile) AddExternalThreadInvite(threadId string) (string, error) {
	_, thrd := tcore.Node.Wallet.GetThread(threadId)
	if thrd == nil {
		return "", errors.New(fmt.Sprintf("could not find thread: %s", threadId))
	}

	// add it
	addr, key, err := thrd.AddExternalInvite()
	if err != nil {
		return "", err
	}

	// create a structured invite
	username, _ := m.GetUsername()
	invite := ExternalInvite{
		Id:      addr.B58String(),
		Key:     string(key),
		Inviter: username,
	}

	return toJSON(invite)
}

// AcceptExternalThreadInvite notifies the thread of a join
func (m *Mobile) AcceptExternalThreadInvite(id string, key string) (string, error) {
	addr, err := tcore.Node.Wallet.AcceptExternalThreadInvite(id, []byte(key))
	if err != nil {
		return "", err
	}
	return addr.B58String(), nil
}

// RemoveThread call core RemoveDevice
func (m *Mobile) RemoveThread(id string) (string, error) {
	addr, err := tcore.Node.Wallet.RemoveThread(id)
	if err != nil {
		return "", err
	}
	return addr.B58String(), err
}

// Devices lists all devices
func (m *Mobile) Devices() (string, error) {
	devices := Devices{Items: make([]Device, 0)}
	for _, dev := range tcore.Node.Wallet.Devices() {
		item := Device{Id: dev.Id, Name: dev.Name}
		devices.Items = append(devices.Items, item)
	}
	return toJSON(devices)
}

// AddDevice calls core AddDevice
func (m *Mobile) AddDevice(name string, pubKey string) error {
	pkb, err := libp2pc.ConfigDecodeKey(pubKey)
	if err != nil {
		return err
	}
	pk, err := libp2pc.UnmarshalPublicKey(pkb)
	if err != nil {
		return err
	}
	return tcore.Node.Wallet.AddDevice(name, pk)
}

// RemoveDevice call core RemoveDevice
func (m *Mobile) RemoveDevice(id string) error {
	return tcore.Node.Wallet.RemoveDevice(id)
}

// AddPhoto adds a photo by path
func (m *Mobile) AddPhoto(path string) (string, error) {
	added, err := tcore.Node.Wallet.AddPhoto(path)
	if err != nil {
		return "", err
	}
	return toJSON(*added)
}

// SharePhoto adds an existing photo to a new thread
func (m *Mobile) AddPhotoToThread(dataId string, key string, threadId string, caption string) (string, error) {
	_, thrd := tcore.Node.Wallet.GetThread(threadId)
	if thrd == nil {
		return "", errors.New(fmt.Sprintf("could not find thread %s", threadId))
	}

	addr, err := thrd.AddPhoto(dataId, caption, []byte(key))
	if err != nil {
		return "", err
	}

	return addr.B58String(), nil
}

// SharePhoto adds an existing photo to a new thread
func (m *Mobile) SharePhotoToThread(dataId string, threadId string, caption string) (string, error) {
	block, err := tcore.Node.Wallet.GetBlockByDataId(dataId)
	if err != nil {
		return "", err
	}
	_, fromThread := tcore.Node.Wallet.GetThread(block.ThreadId)
	if fromThread == nil {
		return "", errors.New(fmt.Sprintf("could not find thread %s", block.ThreadId))
	}
	_, toThread := tcore.Node.Wallet.GetThread(threadId)
	if toThread == nil {
		return "", errors.New(fmt.Sprintf("could not find thread %s", threadId))
	}
	key, err := fromThread.Decrypt(block.DataKeyCipher)
	if err != nil {
		return "", err
	}

	// TODO: owner challenge
	addr, err := toThread.AddPhoto(dataId, caption, key)
	if err != nil {
		return "", err
	}

	return addr.B58String(), nil
}

// GetPhotos returns thread photo blocks with json encoding
func (m *Mobile) GetPhotos(offsetId string, limit int, threadId string) (string, error) {
	_, thrd := tcore.Node.Wallet.GetThread(threadId)
	if thrd == nil {
		return "", errors.New(fmt.Sprintf("thread not found: %s", threadId))
	}

	// build json
	photos := &Photos{Items: make([]Photo, 0)}
	for _, b := range thrd.Blocks(offsetId, limit, repo.PhotoBlock) {
		key, err := thrd.Decrypt(b.DataKeyCipher)
		if err != nil {
			continue
		}
		caption, err := crypto.DecryptAES(b.DataCaptionCipher, key)
		if err != nil {
			continue
		}
		authorId, err := util.IdFromEncodedPublicKey(b.AuthorPk)
		if err != nil {
			continue
		}
		photos.Items = append(photos.Items, Photo{
			Id:       b.DataId,
			Date:     b.Date,
			Caption:  string(caption),
			AuthorId: authorId.Pretty(),
		})
	}
	return toJSON(photos)
}

// GetPhotoData calls GetBlockDataBase64 on a thread
func (m *Mobile) GetPhotoData(id string, path string) (string, error) {
	block, err := tcore.Node.Wallet.GetBlockByDataId(id)
	if err != nil {
		log.Errorf("could not find block for data id %s: %s", id, err)
		return "", err
	}
	_, thrd := tcore.Node.Wallet.GetThread(block.ThreadId)
	if thrd == nil {
		err := errors.New(fmt.Sprintf("could not find thread: %s", block.ThreadId))
		log.Error(err.Error())
		return "", err
	}
	return thrd.GetBlockDataBase64(fmt.Sprintf("%s/%s", id, path), block)
}

// subscribe to thread and pass updates to messenger
func (m *Mobile) subscribe(thrd *thread.Thread) {
	for {
		select {
		case update, ok := <-thrd.Updates():
			if !ok {
				return
			}
			m.messenger.Notify(newEvent("onThreadUpdate", map[string]interface{}{
				"index":       update.Index,
				"thread_id":   update.ThreadId,
				"thread_name": update.ThreadName,
			}))
		}
	}
}

// toJSON returns a json string and logs errors
func toJSON(any interface{}) (string, error) {
	jsonb, err := json.Marshal(any)
	if err != nil {
		log.Errorf("error marshaling json: %s", err)
		return "", err
	}
	return string(jsonb), nil
}
