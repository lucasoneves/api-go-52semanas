package controllers

import (
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
	json.NewEncoder(w).Encode(models.Player)
}

func PlayerDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, player := range models.Player {
		if player.Id == id {
			json.NewEncoder(w).Encode(player)
		}
	}
}
