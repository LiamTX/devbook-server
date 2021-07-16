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
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func openDataseConnection() (*repositories.Users, *sql.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, nil, err
	}

	repository := repositories.NewUserRepository(db)
	return repository, db, nil
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

	if err = user.Prepare(); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	repository, db, err := openDataseConnection()
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

// Find all users by name/nick
func Find(w http.ResponseWriter, r *http.Request) {
	params := strings.ToLower(r.URL.Query().Get("user"))

	repository, db, err := openDataseConnection()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	users, err := repository.FindAll(params)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// Find user by id
func FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	repository, db, err := openDataseConnection()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	user, err := repository.FindOne(userId)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// Update user
func Update(w http.ResponseWriter, r *http.Request) {

}

// Delete user by id
func Delete(w http.ResponseWriter, r *http.Request) {

}
