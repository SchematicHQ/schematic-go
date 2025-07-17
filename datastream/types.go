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
	Error string `json:"error"`
}

type DataStreamClientOptions struct {
	ApiKey         string
	BaseURL        string
	FlagCache      cache.CacheProvider[*rulesengine.Flag]
	CompanyCache   cache.CacheProvider[*rulesengine.Company]
	UserCache      cache.CacheProvider[*rulesengine.User]
	Logger         core.Logger
	MonitorChannel chan bool
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

	// Locks
	flagsMu          sync.RWMutex // For flags cache operations
	companyMu        sync.RWMutex // For company cache operations
	userMu           sync.RWMutex // For user cache operations
	pendingCompReqMu sync.Mutex   // For pending company request operations
	pendingUserReqMu sync.Mutex   // For pending user request operations
	pendingFlagReqMu sync.Mutex   // For pending flag request operations
	writeMu          sync.Mutex   // Existing mutex for WebSocket writes
}

type MessageType string

const (
	MessageTypFull     MessageType = "full"
	MessageTypePartial MessageType = "partial"
	MessageTypeDelete  MessageType = "delete"
	MessageTypeError   MessageType = "error"
)
