package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jscoderdev/npsserverapi/db"
	"github.com/jscoderdev/npsserverapi/models"
	"github.com/jscoderdev/npsserverapi/routes"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HandlerIndex)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	http.ListenAndServe(":4000", r)

}
