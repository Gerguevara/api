package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"api/handlers"
)

func main() {
	router := mux.NewRouter()

	// Index route
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, world!"))
	// })
	router.HandleFunc("/", handlers.HomeHandler)

	http.ListenAndServe(":8080", router)
}
