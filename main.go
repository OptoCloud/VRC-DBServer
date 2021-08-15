package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type JsonConfig struct {
	HttpPort  uint16 `json:"http_port"`
	WebSocket uint16 `json:"websocket_port"`
	MongoDB   struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mongodb"`
}

func ReadConfig() (*JsonConfig, error) {
	configFile, err := os.Open("config.json")
	if err != nil {
		return nil, errors.New("Failed to open file: " + err.Error())
	}
	defer configFile.Close()

	configBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, errors.New("Failed to read file: " + err.Error())
	}

	var config JsonConfig
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, errors.New("Failed to parse file: " + err.Error())
	}

	if len(config.MongoDB.Host) == 0 {
		return nil, errors.New("config does not contain mongodb.host")
	}

	if len(config.MongoDB.Username) == 0 {
		return nil, errors.New("config does not contain mongodb.username")
	}

	if len(config.MongoDB.Password) == 0 {
		return nil, errors.New("config does not contain mongodb.password")
	}

	return &config, nil
}

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal("Failed to open json: ", err)
	}

	errorChain := alice.New(httpLoggerHandler, httpAuthenticatorHandler, httpRecoverHandler)

	var router = mux.NewRouter()
	router.HandleFunc("/user", httpUserPostHandler).Methods("POST")                  // Submit user
	router.HandleFunc("/user/{id}", httpUserGetHandler).Methods("GET")               // Lookup user
	router.HandleFunc("/avatar", httpAvatarPostHandler).Methods("POST")              // Submit avatar
	router.HandleFunc("/avatar/{id}", httpAvatarGetHandler).Methods("GET")           // Lookup avatar
	router.HandleFunc("/world", httpWorldPostHandler).Methods("POST")                // Submit world
	router.HandleFunc("/world/{id}", httpWorldGetHandler).Methods("GET")             // Lookup world
	router.HandleFunc("/search/users", httpSearchUsersGetHandler).Methods("GET")     // Search for users
	router.HandleFunc("/search/avatars", httpSearchAvatarsGetHandler).Methods("GET") // Search for avatars
	router.HandleFunc("/search/worlds", httpSearchWorldsGetHandler).Methods("GET")   // Search for worlds

	http.Handle("/", errorChain.Then(router))

	server := &http.Server{
		Addr: fmt.Sprintf(":%v", config.HttpPort),
	}

	if !dbInit(config.MongoDB.Host, config.MongoDB.Username, config.MongoDB.Password) {
		return
	}
	defer dbClose()

	log.Printf("Server running on port %v\n", config.HttpPort)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
