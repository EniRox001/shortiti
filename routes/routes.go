package routes

import (
	"net/http"

	"github.com/enirox001/shortit/handlers"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Home).Methods("GET")

	return router
}
