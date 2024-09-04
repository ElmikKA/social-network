package main

import (
	"log"
	"net/http"
	"social-network/pkg/api"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: api.Routes(),
	}

	println("starting api server at http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
