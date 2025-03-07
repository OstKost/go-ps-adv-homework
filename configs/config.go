package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
	Auth   AuthConfig
	Email  EmailConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type EmailConfig struct {
	Address  string
	Password string
	SMTPHost string
	SMTPPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		DB: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		Email: EmailConfig{
			Address:  os.Getenv("EMAIL_ADDRESS"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			SMTPHost: os.Getenv("EMAIL_SMTP_HOST"),
			SMTPPort: os.Getenv("EMAIL_SMTP_PORT"),
		},
		Server: ServerConfig{
			Host: getEnvWithDefault("HOST", "localhost"),
			Port: getEnvWithDefault("PORT", "8081"),
		},
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
