package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	responseData := make(map[string]interface{})

	credentials := models.Users{
		Name:        r.FormValue("name"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password"),
		FirstName:   r.FormValue("firstName"),
		LastName:    r.FormValue("lastName"),
		DateOfBirth: r.FormValue("dateOfBirth"),
		Nickname:    r.FormValue("nickname"),
		AboutMe:     r.FormValue("aboutMe"),
	}

	fmt.Println(credentials)

	// check if user is already registered
	exists, err := h.store.CheckUserExists(credentials)
	if err != nil {
		fmt.Println("error checking existing user", err)
		responseData["response"] = "failure"
		responseData["message"] = "database failure"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	if exists {
		fmt.Println("user alreday registered")
		responseData["response"] = "failure"
		responseData["message"] = "Username or email already registered"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	// add user to db
	filePath, err := utils.SaveFile(r, credentials.Name, "Avatar")
	if err != nil {
		fmt.Println("error saving file")
	}
	fmt.Println(filePath)
	credentials.Avatar = filePath

	err = h.store.AddUser(credentials)
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
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

	// check if credentials match
	loggedIn, id, err := h.store.CheckLogin(credentials)
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
	fmt.Println("logged in with :", id)

	// sessions
	err = h.AddSession(w, r, id)
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

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// gets only the logged in user

	responseData := make(map[string]interface{})

	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		fmt.Println("getUser handler err", err)
		responseData["respone"] = "failure"
		responseData["message"] = "Couldn't get user from cookie"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "Getuser successful"
	responseData["getUser"] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
	fmt.Println("sent the response")
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	// fmt.Println("getAllUsers")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// gets all users

	responseData := make(map[string]interface{})

	users, err := h.store.GetAllUsers()
	if err != nil {
		fmt.Println("getAllUsers handler err", err)
		responseData["respone"] = "failure"
		responseData["message"] = "Couldn't get all users"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "GetAllUsers successful"
	responseData["getAllUsers"] = users
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}
