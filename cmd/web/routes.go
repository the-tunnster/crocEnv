package main

import (
	"crocStuff/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler{
	mux := chi.NewRouter()

	mux.Post("/sendFile", handlers.ClientPushing)
	mux.Post("/receiveFile", handlers.ClientPulling)

	mux.Post("/testing", handlers.API_Endpoint_Test)

	return mux
}
