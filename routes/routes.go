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
	r.HandleFunc("/api/players", controllers.ShowAllPlayers).Methods("Get")
	r.HandleFunc("/api/players/{id}", controllers.PlayerDetail).Methods("Get")
	r.HandleFunc("/api/players", controllers.CriarPlayer).Methods("Post")
	r.HandleFunc("/api/players/{id}", controllers.DeletePlayer).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}
