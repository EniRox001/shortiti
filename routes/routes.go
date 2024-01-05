package routes

import (
	"net/http"

	"github.com/enirox001/shortit/handlers"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/", handlers.Shorten).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("GET")

	// public routes
	router.HandleFunc("/public/links/{url}", handlers.LinkDashboard).Methods("GET")

	router.HandleFunc("/{urlId}", handlers.Redirect).Methods("GET")
	router.HandleFunc("/{urlId}/delete", handlers.DeletePublicLink).Methods("GET")
	router.HandleFunc("{urlRoute}", handlers.Redirect).Methods("GET")

	router.HandleFunc("/public/links/{urlId}/qr", handlers.ShowQRCode).Methods("GET")

	return router
}
