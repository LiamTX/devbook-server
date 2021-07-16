package models

import (
	"errors"
	"strings"
	"time"
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

	user.format()

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
	if leg == "create" && user.Password == "" {
		return errors.New("empty_password")
	}

	return nil
}

// Format empty string spaces with TrimSpace
func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
