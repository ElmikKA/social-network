package main

import (
	"fmt"
	"log"
	"net/http"
	"social-network/db"
	"social-network/pkg/api"
)

func main() {

	err := db.InitDb()
	if err != nil {
		fmt.Println("error", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: api.Routes(),
	}

	println("starting api server at http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
