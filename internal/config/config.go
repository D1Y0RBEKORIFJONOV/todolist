package config

import (
	"os"
	"time"
)

type Config struct {
	APP         string
	Environment string
	LogLevel    string
	HostUrl     string
	RedisUrl    string

	Context struct {
		Timeout string
	}
	Token struct {
		Secret    string
		AccessTTL time.Duration
	}
	Email struct {
		From     string
		Password string
		SmtHost  string
		SmtPort  string
	}
	Mongo struct {
		Host           string
		Port           string
		CollectionName string
		DatabaseName   string
	}
	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SslMode  string
	}
}

func Token() string {
	c := Config{}
	c.Token.Secret = getEnv("TOKEN_SECRET", "token_secret")
	return c.Token.Secret
}

func New() *Config {
	var config Config

	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "local")
	config.HostUrl = getEnv("HOST_URL", "todolist_application:9000")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	config.DB.Host = getEnv("POSTGRES_HOST", "postgres")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "+_+diyor2005+_+")
	config.DB.SslMode = getEnv("POSTGRES_SSLMODE", "disable")
	config.DB.Name = getEnv("POSTGRES_DATABASE", "todolist")

	config.Mongo.Host = getEnv("POSTGRES_MONGO_HOST", "mongo")
	config.Mongo.Port = getEnv("POSTGRES_MONGO_PORT", ":27017")
	config.Mongo.DatabaseName = getEnv("MONGO_DBNAME", "todolist")
	config.Mongo.CollectionName = getEnv("COLLECTION_NAME", "details")

	config.Email.SmtHost = getEnv("SMT_HOST", "smtp.gmail.com")
	config.Email.SmtPort = getEnv("SMTP_PORT", "587")
	config.Email.From = getEnv("EMAIL_FROM", "diyordev3@gmail.com")
	config.Email.Password = getEnv("EMAIL_PASSWORD", "ueus bord hbep ttam")

	config.Token.Secret = getEnv("TOKEN_SECRET", "D1YORTOP4EEK")
	accessTTl, err := time.ParseDuration(getEnv("TOKEN_ACCESS_TTL", "1h"))
	if err != nil {
		return nil
	}
	config.Token.AccessTTL = accessTTl
	config.RedisUrl = getEnv("REDIS_URL", "redis:6379")

	return &config
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
