package db

import (
	"database/sql"
	"log"
	"time"

	"gin-example/config"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

// DB is a global database connection pool
var DB *sql.DB

// Connect initializes the database connection pool.
func Connect(cfg config.Config) {
	dsn := config.GetDSN(cfg)
	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %v", err)
	}

	// Setting a connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("Database connected successfully")
}

// Close closes the database connection pool.
func Close() {
	if err := DB.Close(); err != nil {
		log.Fatalf("error closing database connection: %v", err)
	}
}
