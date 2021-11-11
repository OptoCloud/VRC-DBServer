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

	middlewareChain := alice.New(middlewares.Logger, middlewares.Recover)

	var mainRouter = mux.NewRouter()
	authRouter := mainRouter.NewRoute().Subrouter()
	authRouter.Use(middlewares.Authenticator)

	// No auth required to call these
	mainRouter.HandleFunc("/config", handlers.ConfigGet).Methods("GET") // Get API config

	// Websocket stuff
	authRouter.HandleFunc("/ws", wsServer.Upgrade)  // Webscoket endpoint
	authRouter.HandleFunc("/ws/", wsServer.Upgrade) // Webscoket endpoint

	http.Handle("/", middlewareChain.Then(mainRouter))
}
