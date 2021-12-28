package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Port          int    `json:"port"`
	DSN           string `json:"dsn"`
	CronSpec      string `json:"cronSpec"`
	FileUploadDir string `json:"fileUploadDir"`
	SmtpHost      string `json:"smtpHost"`
	SmtpPort      int    `json:"smtpPort"`
	SmtpUsername  string `json:"smtpUsername"`
	SMTPPassword  string `json:"smtpPassword"`
	SmtpTo        string `json:"smtpTo"`
	SmtpSubject   string `json:"smtpSubject"`
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
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(jsonFile)
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
