package configs

import (
	"errors"
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

var Environment = loadConfig()

func loadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load environment file.")
	}

	return Config{
		Host:             getEnvironmentVariableStr("HOST", "localhost"),
		Port:             getEnvironmentVariableStr("PORT", "8080"),
		JWTSecret:        getEnvironmentVariableStr("PORT", "SECRET"),
		SqliteDB:         getEnvironmentVariableStr("SQLITE_DB", "bin/test.db"),
		postgresHost:     getEnvironmentVariableStr("POSTGRES_HOST", "postgres"),
		postgresPort:     getEnvironmentVariableStr("POSTGRES_PORT", "5432"),
		postgresUser:     getEnvironmentVariableStr("POSTGRES_USER", "postgres"),
		postgresPassword: getEnvironmentVariableStr("POSTGRES_PASSWORD", "postgres"),
		postgresDatabase: getEnvironmentVariableStr("POSTGRES_DB", "taskmanager"),
	}
}

func getEnvironmentVariableStr(key string, fallbackValue string) string {
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
