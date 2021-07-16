package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User type
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare call validate and format users methods
func (user *User) Prepare(leg string) error {
	if err := user.validate(leg); err != nil {
		return err
	}

	err := user.format(leg)
	if err != nil {
		return err
	}

	return nil
}

// Validate if is empty field
func (user *User) validate(leg string) error {
	if user.Name == "" {
		return errors.New("empty_name")
	}

	if user.Nick == "" {
		return errors.New("empty_nick")
	}

	if user.Email == "" {
		return errors.New("empty_email")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email_not_valid")
	}

	if leg == "create" && user.Password == "" {
		return errors.New("empty_password")
	}

	return nil
}

// Format empty string spaces with TrimSpace
func (user *User) format(leg string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if leg == "create" {
		hashPassword, err := security.Encrypt(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashPassword)
	}

	return nil
}
