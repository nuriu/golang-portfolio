package configs

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host             string
	Port             string
	JWTSecret        string
	SqliteDB         string
	postgresHost     string
	postgresPort     string
	postgresUser     string
	postgresPassword string
	postgresDatabase string
}

var Environment Config

func init() {
	var err error
	Environment, err = LoadConfig()
	if err != nil {
		log.Fatalf("Error when loading config: %v", err)
	}
}

func loadEnv(envFileName string) error {
	if envFileName == "" {
		envFileName = "../.env"
	}

	err := godotenv.Load(envFileName)
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	return nil

}

func LoadConfig() (Config, error) {
	err := loadEnv("../.env")
	if err != nil {
		return Config{}, fmt.Errorf("failed to load environment: %w", err)
	}

	return Config{
		Host:             GetEnvironmentVariableStr("HOST", "localhost"),
		Port:             GetEnvironmentVariableStr("PORT", "8080"),
		JWTSecret:        GetEnvironmentVariableStr("JWT_SECRET", "SECRET"),
		SqliteDB:         GetEnvironmentVariableStr("SQLITE_DB", "bin/test.db"),
		postgresHost:     GetEnvironmentVariableStr("POSTGRES_HOST", "postgres"),
		postgresPort:     GetEnvironmentVariableStr("POSTGRES_PORT", "5432"),
		postgresUser:     GetEnvironmentVariableStr("POSTGRES_USER", "postgres"),
		postgresPassword: GetEnvironmentVariableStr("POSTGRES_PASSWORD", "postgres"),
		postgresDatabase: GetEnvironmentVariableStr("POSTGRES_DB", "taskmanager"),
	}, nil
}

func GetEnvironmentVariableStr(key string, fallbackValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallbackValue
}

func (config *Config) GetPostgresDsn() (string, error) {
	if config == nil {
		return "", errors.New("config is not loaded, load config before using this method")
	}

	dsn := "host=" + config.postgresHost +
		" user=" + config.postgresUser +
		" password=" + config.postgresPassword +
		" dbname=" + config.postgresDatabase +
		" port=" + config.postgresPort +
		" sslmode=disable TimeZone=UTC"

	return dsn, nil
}
