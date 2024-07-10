package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
}

var Environment = loadConfig()

func loadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load environment file.")
	}

	return Config{
		Host: getEnvironmentVariableStr("HOST", "localhost"),
		Port: getEnvironmentVariableStr("PORT", "8080"),
	}
}

func getEnvironmentVariableStr(key string, fallbackValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallbackValue
}
