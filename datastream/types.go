package datastream

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/schematichq/rulesengine"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/cache"
)

type CompanyCacheProvider cache.CacheProvider[*rulesengine.Company]
type FlagCacheProvider cache.CacheProvider[*rulesengine.Flag]
type UserCacheProvider cache.CacheProvider[*rulesengine.User]

type DataStreamReq struct {
	Action     Action            `json:"action" binding:"required,oneof=start stop"`
	EntityType EntityType        `json:"entity_type" binding:"required,oneof=company flags user"`
	Keys       map[string]string `json:"keys,omitempty"`
}

type DataStreamBaseReq struct {
	Data DataStreamReq `json:"data"`
}

type DataStreamResp struct {
	Data       json.RawMessage `json:"data"`
	EntityID   *string         `json:"entity_id"`
	EntityType string          `json:"entity_type"`
}

type DataStreamClient struct {
	conn                 *websocket.Conn
	logger               schematicgo.Logger
	done                 chan bool
	reconnect            chan bool
	companyCacheProvider CompanyCacheProvider
	flagsCacheProvider   FlagCacheProvider
	userCacheProvider    UserCacheProvider
	companyCache         map[string]*rulesengine.Company
	url                  string
	apiKey               string
}
