package database

import (
	"api/src/config"
	"database/sql"

	// If dont use this import add "_"
	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect and returns database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DB_URI)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
