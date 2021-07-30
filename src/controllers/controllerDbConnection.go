package controllers

import (
	"api/src/database"
	"api/src/repositories"
	"database/sql"
)

func OpenDataseConnection() (*repositories.Users, *sql.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, nil, err
	}

	repository := repositories.NewUserRepository(db)
	return repository, db, nil
}
