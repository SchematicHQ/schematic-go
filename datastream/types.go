package datastream

import (
	"encoding/json"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

type CompanyCacheProvider cache.CacheProvider[*rulesengine.Company]
type FlagCacheProvider cache.CacheProvider[*rulesengine.Flag]
type UserCacheProvider cache.CacheProvider[*rulesengine.User]

type DataStreamReq struct {
	Action     Action            `json:"action"`
	EntityType EntityType        `json:"entity_type"`
	Keys       map[string]string `json:"keys,omitempty"`
}

type DataStreamBaseReq struct {
	Data DataStreamReq `json:"data"`
}

type DataStreamResp struct {
	Data        json.RawMessage `json:"data"`
	EntityID    *string         `json:"entity_id"`
	EntityType  string          `json:"entity_type"`
	MessageType MessageType     `json:"message_type"`
}

type DataStreamError struct {
	Code  int32  `json:"code"`
	Error string `json:"error"`
}

type DataStreamClient struct {
	cacheTTL             time.Duration
	conn                 *websocket.Conn
	logger               core.Logger
	done                 chan bool
	ctxErrors            chan *core.CtxError
	reconnect            chan bool
	monitorChannel       chan bool
	companyCacheProvider CompanyCacheProvider
	flagsCacheProvider   FlagCacheProvider
	userCacheProvider    UserCacheProvider
	companyCache         map[string]*rulesengine.Company
	url                  *url.URL
	apiKey               string

	pendingCompanyRequests map[string][]chan *rulesengine.Company
	pendingUserRequests    map[string][]chan *rulesengine.User
	pendingFlagRequest     chan bool
	mu                     sync.RWMutex
}

type MessageType string

const (
	MessageTypFull     MessageType = "full"
	MessageTypePartial MessageType = "partial"
	MessageTypeDelete  MessageType = "delete"
)
