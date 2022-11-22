package main

import (
	"net/http"
	"user-service/handlers"
)

func (app *Config) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/create-user", handlers.CreateUserHandler)

	return mux
}
