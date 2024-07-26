package main

import (
	"task-manager/configs"
	_ "task-manager/docs"
	"task-manager/internal/server"

	gormSqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title REST API
// @version 1.0
// @description REST API documentation.
// @host localhost:8080
func main() {
	db, err := gorm.Open(gormSqlite.Open(configs.Environment.SqliteDB), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	server.Run(configs.Environment.Host+":"+configs.Environment.Port, db)
}
