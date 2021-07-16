package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func openDataseConnection() (*repositories.Users, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	repository := repositories.NewUserRepository(db)
	return repository, nil
}

// Create user
func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	repository, err := openDataseConnection()
	if err != nil {
		log.Fatal(err)
	}

	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("user id: %d", userId)))
}

// Find all users
func Find(w http.ResponseWriter, r *http.Request) {

}

// Find user
func FindOne(w http.ResponseWriter, r *http.Request) {

}

// Update user
func Update(w http.ResponseWriter, r *http.Request) {

}

// Delete user
func Delete(w http.ResponseWriter, r *http.Request) {

}
