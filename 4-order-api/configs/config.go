package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB   DBConfig
	Sms  SmsConfig
	Auth AuthConfig
}

type DBConfig struct {
	Dsn string
}

type SmsConfig struct {
	ApiId string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		DB: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
		Sms: SmsConfig{
			ApiId: os.Getenv("SMS_RU_ID"),
		},
	}
}
