package config

import (
	"hayday/server/src/models"
	"sync"

	_websocket "nhooyr.io/websocket"
)

const (
	DB_SET_MAX_OPEN_CONNS = 10
	DB_SET_MAX_IDLE_CONNS = 10
	CONFIG_FILEPATH       = "."
	CONFIG_FILENAME       = "config"
	CONFIG_FILE_EXTENSION = "json"
	DYNAMIC_LINK          = "https://firebasedynamiclinks.googleapis.com/v1/shortLinks"
	DOMAIN_URI_PREFIX     = "https://getstan.page.link"
	ANDROID_PACKAGE_NAME  = "com.getstan"
	IOS_BUNDLE_ID         = "com.getstan"
)

var (
	FIRESTORE_CLUB_COLL                = "PROD_CLUB_ROOMS"
	FIRESTORE_CREATOR_WAITLISTING_COLL = "PROD_CREATOR_WAITLISTING"
	EnvValues                          serverConfig
	WsClients                          = make(map[string]map[*_websocket.Conn]models.WebsocketModel)
	WsClientMux                        sync.Mutex
	DefaultClubAdmins                  = []int32{167824, 945613, 16283, 408503, 5283105}

	RedisPubSubConsumers = make(map[string]bool)
	RedisStreamMap       = make(map[string]map[int32]map[string]interface{})
	ClubWiseMessages     = make(map[string][]interface{})
)
