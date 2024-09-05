package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	responseData := make(map[string]interface{})

	credentials := models.RegisterCredentials{}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	data, _ := json.MarshalIndent(credentials, "", "  ")
	fmt.Println(string(data))

	// check if user is already registered

	// add user to db

	responseData["response"] = "success"
	responseData["message"] = "User registered successfully"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseData)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	responseData := make(map[string]interface{})

	credentials := models.LoginCredentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Println(credentials)

	// check if credentials match

	// sessions
	// removes previous sessions with this user
	// adds new session for user

	responseData["response"] = "success"
	responseData["message"] = "Login successful"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}
