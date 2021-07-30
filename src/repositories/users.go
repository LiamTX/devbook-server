package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Find all users by name/nick
func (u Users) FindAll(param string) ([]models.User, error) {
	param = fmt.Sprintf("%%%s%%", param)

	rows, err := u.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		param, param,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Find user by id
func (u Users) FindOne(id uint64) (models.User, error) {
	rows, err := u.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update user by id
func (u Users) UpdateOne(id uint64, user models.User) error {
	statement, err := u.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		id,
	); err != nil {
		return err
	}

	return nil
}

// Delete user by id
func (u Users) DeleteOne(id uint64) error {
	statement, err := u.db.Prepare(
		"delete from users where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

// Find one user by email and returns id and password hashed
func (u Users) FindByEmail(email string) (models.User, error) {
	row, err := u.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
