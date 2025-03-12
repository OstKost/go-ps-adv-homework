package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
	Sms    SmsConfig
	Auth   AuthConfig
}

type ServerConfig struct {
	Host string
	Port string
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
			Secret: os.Getenv("TOKEN"),
		},
		Sms: SmsConfig{
			ApiId: os.Getenv("SMS_RU_ID"),
		},
		Server: ServerConfig{
			Host: getEnvWithDefault("HOST", "localhost"),
			Port: getEnvWithDefault("PORT", "8081"),
		},
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
