package httpServer

import (
	"fmt"
	"log"
	"net/http"
)

var server *http.Server

func Run(port uint16) error {
	server = &http.Server{
		Addr: fmt.Sprintf("localhost:%v", port),
	}
	log.Printf("Server running on port %v...\n", port)
	return server.ListenAndServe()
}
func Close() error {
	return server.Close()
}
