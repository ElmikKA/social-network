package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/db"
	"social-network/pkg/models"
	"social-network/pkg/sessions"
)

func (api *APIServer) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	responseData := make(map[string]interface{})

	credentials := models.Users{}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println("Register error decoding json ", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	data, _ := json.MarshalIndent(credentials, "", "  ")
	fmt.Println(string(data))

	// check if user is already registered
	exists, err := db.CheckUserExists(api.db, credentials)
	if err != nil {
		fmt.Println("error checking existing user", err)
		responseData["response"] = "failure"
		responseData["message"] = "database failure"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	if exists {
		responseData["response"] = "failure"
		responseData["message"] = "Username or email already registered"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	// add user to db
	err = db.AddUser(api.db, credentials)
	if err != nil {
		fmt.Println("register error adding new user", err)
		responseData["response"] = "failure"
		responseData["message"] = "db error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
	}

	responseData["response"] = "success"
	responseData["message"] = "User registered successfully"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseData)
}

func (api *APIServer) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	responseData := make(map[string]interface{})

	credentials := models.LoginCredentials{}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println("error", err)
		responseData["response"] = "failure"
		responseData["message"] = "invalid JSON payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	fmt.Println(credentials)

	// check if credentials match
	loggedIn, id, err := db.CheckLogin(api.db, credentials)
	if err != nil {
		fmt.Println("login error CheckLogin", err)
		responseData["response"] = "failure"
		responseData["message"] = "db error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	if !loggedIn {
		fmt.Println("invalid credentials")
		responseData["response"] = "failure"
		responseData["message"] = "invalid credentials"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	fmt.Println("logged in with id:", id)

	// sessions
	err = sessions.AddSession(w, r, api.db, id)
	if err != nil {
		fmt.Println("error adding session", err)
		responseData["response"] = "failure"
		responseData["message"] = "failed to add a new session"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "Login successful"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}

func (api *APIServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getUsers")
}
