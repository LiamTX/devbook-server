package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func openDataseConnection() (*repositories.Users, error, *sql.DB) {
	db, err := database.Connect()
	if err != nil {
		return nil, err, nil
	}

	repository := repositories.NewUserRepository(db)
	return repository, nil, db
}

// Create user
func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	repository, err, db := openDataseConnection()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	user.ID, err = repository.Create(user)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
	return
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
