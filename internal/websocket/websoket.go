package websocket

import (
	"hayday/server/config"
	"time"

	"hayday/server/internal/logger"
	"hayday/server/internal/middleware"

	"fmt"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"

	"github.com/google/uuid"
	_websocket "nhooyr.io/websocket"
)

func InitiateSocketServer() {
	apiServerHandler := http.NewServeMux()
	apiServerHandler.HandleFunc("/api/health", heathCheck)

	go func() {
		apiServer := &http.Server{
			Addr:    fmt.Sprintf(":%s", config.EnvValues.SERVER_PORT),
			Handler: apiServerHandler,
		}
		logger.Log.Infof(`Initializing the api server in PORT : %s`, config.EnvValues.SERVER_PORT)
		if err := apiServer.ListenAndServe(); err != nil {
			logger.Log.Fatalf(`failed to initialize the api server.... %s`, err)
		}
	}()

	wsServerHandler := http.NewServeMux()
	wsServerHandler.HandleFunc("/api/v5/health", heathCheck)
	//apiServerHandler.HandleFunc("/api/v5/club", middleware.WsAuthValidator(http.HandlerFunc(createClub.CreateClub)))
	handlerChain := http.HandlerFunc(WsHandler)                   // 3
	handlerChain = middleware.WsAuthValidator(handlerChain)       // 2
	handlerChain = middleware.WsConnectionValidator(handlerChain) // 1
	wsServerHandler.HandleFunc("/ws/club", handlerChain)

	// <----- pprof endpoints ----->
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/", pprof.Index)
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/cmdline", pprof.Cmdline)
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/profile", pprof.Profile)
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/symbol", pprof.Symbol)
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/trace", pprof.Trace)
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
	wsServerHandler.HandleFunc("/ws/club/api/v5/debug/pprof/block", pprof.Handler("block").ServeHTTP)

	go func() {
		wsServer := &http.Server{
			Addr:         fmt.Sprintf(":%s", config.EnvValues.WEBSOCKET_SERVER_PORT),
			Handler:      wsServerHandler,
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}
		logger.Log.Infof(`Initializing the Websocket server in PORT : %s`, config.EnvValues.WEBSOCKET_SERVER_PORT)
		if err := wsServer.ListenAndServe(); err != nil {
			logger.Log.Fatalf(`failed to initialize the websocket server.... %s`, err)
		}
	}()

	select {}
}

func heathCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Fuck Yeah!")
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	requestId, _ := uuid.NewUUID()
	log := logger.Log.WithFields(map[string]interface{}{
		"EVENT":      "WEBSOCKET_HANDLER",
		"REQUEST_ID": requestId,
	})
	queryValues := r.URL.Query()
	userId := queryValues.Get("userId")
	clubId := queryValues.Get("clubId")

	wsConn, err := _websocket.Accept(w, r, nil)
	if err != nil {
		log.Errorf("Error accepting WebSocket connection: %v", err)
		return
	}

	wsConn.SetReadLimit(-1)
	//utils.AddSocketConnectionToMap(wsConn, userId, clubId, log)
	log.Infof("Client connected to websocket: CLUB_ID: %s, USER_ID: %d, TOTAL CONNECTIONS: %d", clubId, userId, len(config.WsClients[clubId]))

	// redisHelper.InitializeConsumers(log, clubId)
	//redisHelper.InitializeConsumers(r.Context(), log, clubId)
	//wsMsgReader.WsReader(r.Context(), log, wsConn, requestId, userId, clubId)

	log.Infof("Client disconnected: CLUB_ID: %s, USER_ID: %d, TOTAL CONNECTIONS: %d", clubId, userId, len(config.WsClients[clubId]))
	//utils.CloseSocketConnection(r.Context(), wsConn, clubId, _websocket.StatusNormalClosure, "closed", log)
}
