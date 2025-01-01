package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB   *sql.DB
	once sync.Once
)

// Initialize connects to the database and ensures schema exists.
func InitializeDatabase() error {
	var err error
	once.Do(func() {
		DB, err = sql.Open("sqlite3", "./blogs.db")
		if err != nil {
			err = fmt.Errorf("connection failed: %v", err)
			return
		}
		if pingErr := DB.Ping(); pingErr != nil {
			err = fmt.Errorf("ping failed: %v", pingErr)
			return
		}
		_, execErr := DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			author TEXT NOT NULL,
			timestamp TEXT NOT NULL
		);`)
		if execErr != nil {
			err = fmt.Errorf("table creation failed: %v", execErr)
			return
		}
		log.Println("Database connected and schema ensured.")
	})
	return err
}

// GetDB returns the database instance.
func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database not initialized.")
	}
	return DB
}

// Close shuts the database connection.
func CloseDatabase() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}
