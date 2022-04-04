package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	log.Println("Starting server at port ", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
