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
	JwtKey        string `json:"jwtKey"`
	CronSpec      string `json:"cronSpec"`
	FileUploadDir string `json:"fileUploadDir"`
	SmtpHost      string `json:"smtpHost"`
	SmtpPort      int    `json:"smtpPort"`
	SmtpUsername  string `json:"smtpUsername"`
	SMTPPassword  string `json:"smtpPassword"`
	SmtpSubject   string `json:"smtpSubject"`
}

var defaultConfig = Config{
	Port:          8060,
	DSN:           "host=localhost user=my_flomo password=my_flomo dbname=my_flomo port=5432 sslmode=disable TimeZone=Asia/Shanghai",
	JwtKey:        "jwT_p@sSw0rd",
	CronSpec:      "0 20 * * *",
	FileUploadDir: "./",
	SmtpHost:      "smtp-mail.outlook.com",
	SmtpPort:      587,
	SmtpUsername:  "",
	SMTPPassword:  "",
}

var Data *Config

func init() {
	config, err := readConfig()
	if err != nil {
		log.Println("read config error:", err)
		log.Println("use default config", defaultConfig)
		config = &defaultConfig
	}
	Data = config
}

func readConfig() (*Config, error) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		_ = jsonFile.Close()
		return &Config{}, err
	}

	jsonFileByte, _ := ioutil.ReadAll(jsonFile)
	_ = jsonFile.Close()

	var config Config
	err = json.Unmarshal(jsonFileByte, &config)
	if err != nil {
		return &Config{}, err
	}

	return &config, nil
}
