package routes

import (
	"api/controllers"
	"api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/players", controllers.ShowAllPlayers).Methods("Get")
	r.HandleFunc("/api/players/{id}", controllers.PlayerDetail).Methods("Get")
	r.HandleFunc("/api/players", controllers.CriarPlayer).Methods("Post")
	r.HandleFunc("/api/players/{id}", controllers.DeletePlayer).Methods("Delete")
	r.HandleFunc("/api/players/{id}", controllers.EditaPlayer).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
