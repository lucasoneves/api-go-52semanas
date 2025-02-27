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
