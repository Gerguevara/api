package handlers

import (
	"api/db"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	w.Header().Set("content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	//toma params del reques
	params := mux.Vars(r)
	//deckara una variable tipo user inicializado con el zero value
	var user models.User
	fmt.Println(params["id"])
	//busca el usuario con el id, y si esta lo asigna al
	db.DB.Find(&user, "id = ?", params["id"])
	fmt.Println(user)
	//si no lo encuentra le responde con el nofounf
	if user.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	//trae las task asociadas a el y como tas es un arreglo
	//las va agregando
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	w.Header().Set("content-Type", "application/json; charset=utf-8")
	// responseUser, _ := user.MarshalJSON()
	// w.Write([]byte(responseUser))
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
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.Header().Set("content-Type", "application/json; charset=utf-8")

	w.WriteHeader(http.StatusOK)
}
