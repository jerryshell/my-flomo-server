package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port         int    `json:"port"`
	DSN          string `json:"dsn"`
	JwtKey       string `json:"jwtKey"`
	CronSpec     string `json:"cronSpec"`
	SmtpHost     string `json:"smtpHost"`
	SmtpPort     int    `json:"smtpPort"`
	SmtpUsername string `json:"smtpUsername"`
	SMTPPassword string `json:"smtpPassword"`
	SmtpSubject  string `json:"smtpSubject"`
}

var Data *Config = &Config{
	Port:         8060,
	DSN:          "host=localhost user=my_flomo password=my_flomo dbname=my_flomo port=5432 sslmode=disable TimeZone=Asia/Shanghai",
	JwtKey:       "jwT_p@sSw0rd",
	CronSpec:     "0 20 * * *",
	SmtpHost:     "smtp-mail.outlook.com",
	SmtpPort:     587,
	SmtpUsername: "",
	SMTPPassword: "",
	SmtpSubject:  "My Flomo 每日回顾",
}

func init() {
	// Port
	port := os.Getenv("PORT")
	if port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			log.Println("port strconv.Atoi :: err", err)
		} else {
			Data.Port = portInt
		}
	}

	// DSN
	dsn := os.Getenv("DSN")
	if dsn != "" {
		Data.DSN = dsn
	}

	// JwtKey
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey != "" {
		Data.JwtKey = jwtKey
	}

	// CronSpec
	cronSpec := os.Getenv("CRON_SPEC")
	if cronSpec != "" {
		Data.CronSpec = cronSpec
	}

	// SmtpHost
	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost != "" {
		Data.SmtpHost = smtpHost
	}

	// SmtpPort
	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort != "" {
		smtpPortInt, err := strconv.Atoi(smtpPort)
		if err != nil {
			log.Println("smtpPort strconv.Atoi :: err", err, "use default smtpPort ::", Data.SmtpPort)
		} else {
			Data.SmtpPort = smtpPortInt
		}
	}

	// SmtpUsername
	smtpUsername := os.Getenv("SMTP_USERNAME")
	if smtpUsername != "" {
		Data.SmtpUsername = smtpUsername
	}

	// SMTPPassword
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	if smtpPassword != "" {
		Data.SMTPPassword = smtpPassword
	}

	// SmtpSubject
	smtpSubject := os.Getenv("SMTP_SUBJECT")
	if smtpSubject != "" {
		Data.SmtpSubject = smtpSubject
	}
}
