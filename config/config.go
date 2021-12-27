package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Version      string `json:"version"`
	DSN          string `json:"dsn"`
	SmtpHost     string `json:"smtpHost"`
	SmtpPort     int    `json:"smtpPort"`
	SmtpUsername string `json:"smtpUsername"`
	SmtpPassword string `json:"smtpPassword"`
	SmtpTo       string `json:"smtpTo"`
	SmtpSubject  string `json:"smtpSubject"`
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
