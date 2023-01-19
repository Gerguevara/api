package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"api/db"
	"api/handlers"
	"api/models"
)

func main() {
	db.Connect()
	// db.DB.Migrator().DropTable(models.User{})
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	// Index route
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, world!"))
	// })
	router.HandleFunc("/", handlers.HomeHandler)
	// users routes
	router.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", handlers.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
