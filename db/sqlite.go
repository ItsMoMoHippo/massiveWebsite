package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var LDB *sql.DB

func LInit() {
	sqlitePath := filepath.Join("secrets", "test.db")

	var err error
	LDB, err = sql.Open("sqlite3", sqlitePath)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	err = LDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("✅ Connected to database successfully.")
}
