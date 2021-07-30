package controllers

import (
	"api/src/models"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login auth user
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	repository, db, err := OpenDataseConnection()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userExists, err := repository.FindByEmail(user.Email)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Decrypt(userExists.Password, user.Password); err != nil {
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}

	responses.JSON(w, http.StatusOK, "Logado!")
}
