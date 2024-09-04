package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := models.Users{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("error", err)
			return
		}

	}

}
