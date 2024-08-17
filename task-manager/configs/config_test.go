package configs_test

import (
	"os"
	"task-manager/configs"
	"testing"
)

func setup() {
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("PORT", "8081")
	os.Setenv("JWT_SECRET", "MOCK_SECRET")
	os.Setenv("SQLITE_DB", "test.db")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5433")
	os.Setenv("POSTGRES_USER", "testuser")
	os.Setenv("POSTGRES_PASSWORD", "testpassword")
	os.Setenv("POSTGRES_DB", "testdb")
}

func TestConfig(t *testing.T) {
	setup()

	config, err := configs.LoadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Run("Environment variables should load correctly", func(t *testing.T) {
		if config.Host != "127.0.0.1" {
			t.Errorf("Expected Host to be '127.0.0.1', got '%s'", config.Host)
		}

		if config.Port != "8081" {
			t.Errorf("Expected Port to be '8081', got '%s'", config.Port)
		}

		if config.JWTSecret != "MOCK_SECRET" {
			t.Errorf("Expected JWTSecret to be 'MOCK_SECRET', got '%s'", config.JWTSecret)
		}

		if config.SqliteDB != "test.db" {
			t.Errorf("Expected SqliteDB to be 'test.db', got '%s'", config.SqliteDB)
		}
	})

	t.Run("should construct correct dsn for postgres connection", func(t *testing.T) {
		dsn, err := config.GetPostgresDsn()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expectedDsn := "host=localhost user=testuser password=testpassword dbname=testdb port=5433 sslmode=disable TimeZone=UTC"
		if dsn != expectedDsn {
			t.Errorf("expected dsn to be '%s', got '%s'", expectedDsn, dsn)
		}
	})

	// Cleanup environment variables after the test
	os.Clearenv()
}

func TestLoadEnvError(t *testing.T) {
	// Temporarily rename or remove the .env file
	err := os.Rename("../.env", "../.env.bak")
	if err != nil {
		t.Fatal("Failed to rename .env file for testing")
	}

	// Restore the .env file after test
	defer os.Rename("../.env.bak", "../.env")

	// Run loadConfig which should now fail
	_, err = configs.LoadConfig()
	if err == nil {
		t.Fatal("Expected error when .env file is missing, got nil")
	}
}

func TestGetPostgresDsnNilConfig(t *testing.T) {
	var config *configs.Config
	_, err := config.GetPostgresDsn()
	if err == nil {
		t.Error("Expected error when Config is nil, got nil")
	}
}

func TestGetEnvironmentVariableStr(t *testing.T) {
	os.Setenv("TEST_VAR", "test_value")
	value := configs.GetEnvironmentVariableStr("TEST_VAR", "default_value")
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	os.Unsetenv("TEST_VAR")
	value = configs.GetEnvironmentVariableStr("TEST_VAR", "default_value")
	if value != "default_value" {
		t.Errorf("Expected 'default_value', got '%s'", value)
	}
}
