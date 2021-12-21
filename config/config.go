package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Version string `json:"version"`
	DSN     string `json:"dsn"`
}

var Data *Config

func init() {
	config, err := readConfig()
	if err != nil {
		panic(err)
	}
	Data = config
}

func readConfig() (*Config, error) {
	jsonFile, err := os.Open("config.json")
	defer jsonFile.Close()
	if err != nil {
		return &Config{}, err
	}

	jsonFileByte, _ := ioutil.ReadAll(jsonFile)
	var config Config
	err = json.Unmarshal(jsonFileByte, &config)
	if err != nil {
		return &Config{}, err
	}

	return &config, nil
}
