package main

import (
	"net/http"
	"user-service/config"
	"user-service/handlers"
)

func routes(app *config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/create-user", handlers.CreateHandlerFunc(app))

	return mux
}
