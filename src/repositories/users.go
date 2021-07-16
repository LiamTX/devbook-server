package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

// Create a user repository
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Insert user into database
func (u Users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		user.Password,
	)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
