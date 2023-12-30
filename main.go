package main

import (
	"net/http"

	"github.com/enirox001/shortit/data"
	"github.com/enirox001/shortit/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handlers := routes.New()

	server := &http.Server{
		Addr:    ":3000",
		Handler: handlers,
	}

	data.ConnectDatabase()

	server.ListenAndServe()
}
