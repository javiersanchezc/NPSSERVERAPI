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
	r.HandleFunc("/GetFiles", routes.GetConvertFiles).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/cardif", routes.GetLoaInsuranceCardif).Methods("GET")

	// Agregar middleware para manejar CORS
	headersMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	// Aplicar middleware CORS a todas las rutas
	r.Use(headersMiddleware)

	http.ListenAndServe(":4000", r)
}
