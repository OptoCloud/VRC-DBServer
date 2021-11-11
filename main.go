package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"vrcdb/httpServer"
)

type JsonConfig struct {
	HttpPort uint16 `json:"http_port"`
	MongoDB  struct {
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

	httpServer.Init()
	if !dbInit(config.MongoDB.Host, config.MongoDB.Username, config.MongoDB.Password) {
		return
	}
	defer dbClose()

	httpServer.Run(config.HttpPort)
}
