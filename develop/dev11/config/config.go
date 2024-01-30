package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ServerHost string `json:"server_host"`
	ServerPort string `json:"server_port"`
}

func InitialzeConfig(path string) Config {
	config := Config{}
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Reading config file error: ", err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Unmarshalsing config file error: ", err)
	}
	return config
}
