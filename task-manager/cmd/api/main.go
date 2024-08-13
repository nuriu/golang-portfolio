package main

import (
	"log"
	"task-manager/configs"
	_ "task-manager/docs"
	"task-manager/internal/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title REST API
// @version 1.0
// @description REST API documentation.
// @host localhost:8080
func main() {
	dsn, err := configs.Environment.GetPostgresDsn()
	if err != nil {
		log.Fatal("Failed to construct postgres dsn:", err)
	}
	// db, err := gorm.Open(gormSqlite.Open(configs.Environment.SqliteDB), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	server.Run(configs.Environment.Host+":"+configs.Environment.Port, db)
}
