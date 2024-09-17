package main

import (
	"fmt"
	"social-network/db"
	"social-network/pkg/api"
)

func main() {
	trigger := false //if you drop database set it to true and run code once. turn if off after
	Db, err := db.InitDb(trigger)
	if err != nil {
		fmt.Println("database init error", err)
	}

	server := api.NewAPIServer(":8080", Db)
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
