package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
	"strconv"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != "POST" {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	credentials := models.Users{
		Name:        r.FormValue("name"),
		Email:       r.FormValue("email"),
		Password:    r.FormValue("password"),
		FirstName:   r.FormValue("firstName"),
		LastName:    r.FormValue("lastName"),
		DateOfBirth: r.FormValue("dateOfBirth"),
		Nickname:    r.FormValue("nickname"),
		AboutMe:     r.FormValue("aboutMe"),
		Privacy:     r.FormValue("privacy"),
	}

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

	filePath, err := utils.SaveFile(r, credentials.Name, "Avatar")
	if err != nil {
		if filePath == "No avatar" {
			fmt.Println("no avatar")
		} else {

			fmt.Println("error saving file")
			responseData["response"] = "failure"
			responseData["message"] = "Internal server error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
		}
	}
	if filePath != "No avatar" {
		credentials.Avatar = filePath
	}

	err = h.store.AddUser(credentials)
	if err != nil {
		fmt.Println("register error adding new user", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
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
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != "POST" {
		fmt.Println("inside not post")
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		json.NewEncoder(w).Encode(responseData)
		w.Header().Set("Content-Type", "application/json")
		return
	}

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

func (h *Handler) LogOut(w http.ResponseWriter, r *http.Request) {

	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != "DELETE" {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		json.NewEncoder(w).Encode(responseData)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		fmt.Println("error getting user from cookie", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		json.NewEncoder(w).Encode(responseData)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	err = h.store.DeleteSession(user.Id)
	if err != nil {
		fmt.Println("error deleting session", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		json.NewEncoder(w).Encode(responseData)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	responseData["response"] = "success"
	responseData["message"] = "User logged out successfully"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != "GET" {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		json.NewEncoder(w).Encode(responseData)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	userId, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		fmt.Println("err getuser url", err)
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	user, err := h.store.GetUser(userId)

	if err != nil {
		fmt.Println("getUser handler err", err)
		responseData["respone"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	// add followers,followings
	contacts, err := h.store.GetContacts(user.Id)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	// posts
	posts, err := h.store.GetAllUserPosts(user.Id)
	if err != nil && err != sql.ErrNoRows {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	if err != sql.ErrNoRows {
		responseData["posts"] = posts
	}

	status, err := h.store.IsFollowing(h.id, userId)
	if err != nil && err != sql.ErrNoRows {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	if h.id == userId {
		responseData["following"] = "self"
		responseData["ownPage"] = true
	} else if err == sql.ErrNoRows {
		responseData["following"] = "not following"
		responseData["ownPage"] = false
	} else {
		responseData["following"] = status
		responseData["ownPage"] = false
	}

	responseData["response"] = "success"
	responseData["message"] = "Getuser successful"
	responseData["getUser"] = user
	responseData["followers"] = contacts

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
	fmt.Println("sent the response")
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != "GET" {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		json.NewEncoder(w).Encode(responseData)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	users, err := h.store.GetAllUsers()
	if err != nil {
		fmt.Println("getAllUsers handler err", err)
		responseData["respone"] = "failure"
		responseData["message"] = "Internal server error"
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

func (h *Handler) CheckLogin(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	responseData["userId"] = h.id
	responseData["response"] = "success"
	responseData["message"] = "logged in"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
