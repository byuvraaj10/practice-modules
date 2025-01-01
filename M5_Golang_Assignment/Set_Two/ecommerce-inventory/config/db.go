package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase() (*sql.DB, error) {
	// Open database connection
	db, err := sql.Open("sqlite3", "./ecommerce.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return nil, err
	}

	// Create users table
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT
	);`); err != nil {
		log.Fatal("Error creating users table: ", err)
		return nil, err
	}

	// Create products table
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		stock INTEGER,
		category_id INTEGER
	);`); err != nil {
		log.Fatal("Error creating products table: ", err)
		return nil, err
	}

	return db, nil
}
