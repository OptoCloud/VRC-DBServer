package httpServer

import (
	"log"
	"net/http"
	"vrcdb/httpServer/handlers"
	"vrcdb/httpServer/middlewares"
	"vrcdb/wsServer"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func Init() {
	log.Println("Initializing http routes...")

	middlewareChain := alice.New(middlewares.Logger, middlewares.ValidateHwid, middlewares.Recover)

	var mainRouter = mux.NewRouter()
	authRouter := mainRouter.NewRoute().Subrouter()
	authRouter.Use(middlewares.Authenticator)

	// No auth required to call these
	mainRouter.HandleFunc("/health", handlers.HealthGet).Methods("GET")              // Get API health
	mainRouter.HandleFunc("/config", handlers.ConfigGet).Methods("GET")              // Get API config
	mainRouter.HandleFunc("/account", handlers.AuthLoginGet).Methods("PUT")       // Get auth cookie with credentials
	authRouter.HandleFunc("/register", handlers.AuthLogoutPut).Methods("PUT")     // Destroy auth cookie

	// Websocket stuff
	authRouter.HandleFunc("/ws", wsServer.Upgrade)  // Webscoket endpoint
	authRouter.HandleFunc("/ws/", wsServer.Upgrade) // Webscoket endpoint

	http.Handle("/", middlewareChain.Then(mainRouter))

	handlers.ServerShutdownFunc = func() { server.Close() }
}
