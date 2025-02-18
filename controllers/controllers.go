package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ShowAllPlayers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Player)
}
