package config

import (
	"log"
	"os"
	"strconv"
	// "github.com/google/s2a-go/fallback"
)

type Config struct {
	AppPort          string
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresSSLMode  string

	RedisAddr     string
	RedisPassword string
	RedisDB       string

	RabbitMQURL      string
	CORSAllowOrigins string

	HoldTTLMinutes int
}

func MustLoad() Config {

	cfg := Config{
		AppPort:          getEnv("APP_PORT", "8080"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresDB:       getEnv("POSTGRES_DB", "concert_booking"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresSSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisDB:          getEnv("REDIS_DB", "0"),
		RabbitMQURL:      getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		CORSAllowOrigins: getEnv("CORS_ALLOW_ORIGINS", "http://localhost:3000"),
		HoldTTLMinutes:   getEnvAsInt("TTL_MINUTE", 5),
	}

	log.Println("config loaded")

	return cfg
}

func getEnv(key, fallback string) string {

	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("invalid int env %s=%q, fallback=%d", key, v, fallback)
		return fallback
	}

	return n
}