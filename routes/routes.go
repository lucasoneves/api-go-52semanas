package routes

import (
	"api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/players", controllers.ShowAllPlayers)
	log.Fatal(http.ListenAndServe(":8000", r))
}
