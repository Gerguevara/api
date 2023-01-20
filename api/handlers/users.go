package handlers

import (
	"api/db"
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	//toma params del reques
	params := mux.Vars(r)
	//deckara una variable tipo user inicializado con el zero value
	var user models.User
	//busca el usuario con el id, y si esta lo asigna al
	db.DB.First(&user, params["id"])

	//si no lo encuentra le responde con el nofounf
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	//trae las task asociadas a el y como tas es un arreglo
	//las va agregando
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	//declara la variable con el zero value
	var user models.User
	//decodifica el body del request y lo asigna al user
	json.NewDecoder(r.Body).Decode(&user)
	//intenta hacer el guardado
	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
