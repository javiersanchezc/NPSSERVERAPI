package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jscoderdev/npsserverapi/db"
	"github.com/jscoderdev/npsserverapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
	w.Write([]byte("Get Users "))

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user) // Convert request body to json and store it in the 'user' variable.
	createdUser := db.DB.Create(&user)    // Save the data into the database.

	json.NewEncoder(w).Encode(createdUser) // Send the created user back as a response with JSON format.
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&user)
	}

}
