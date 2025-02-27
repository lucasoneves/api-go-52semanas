package controllers

import (
	"api/database"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ShowAllPlayers(w http.ResponseWriter, r *http.Request) {
	var p []models.Players
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func PlayerDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var player models.Players
	database.DB.First(&player, id)
	json.NewEncoder(w).Encode(player)
}

func CriarPlayer(w http.ResponseWriter, r *http.Request) {
	var newPlayer models.Players
	json.NewDecoder(r.Body).Decode(&newPlayer)
	database.DB.Create(&newPlayer)
	json.NewEncoder(w).Encode(newPlayer)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var player models.Players
	database.DB.Delete(&player, id)
	json.NewEncoder(w).Encode(player)
}

func EditaPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var player models.Players
	database.DB.First(&player, id)
	json.NewDecoder(r.Body).Decode(&player)
	database.DB.Save(&player)
	json.NewEncoder(w).Encode(player)
}
