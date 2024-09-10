package main

import (
	"fmt"
	"social-network/db"
	"social-network/pkg/api"
)

func main() {
	Db, err := db.InitDb()
	if err != nil {
		fmt.Println("database init error", err)
	}

	server := api.NewAPIServer(":8080", Db)
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
