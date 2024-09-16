package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) AddFollow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adding follower")
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

	var data struct {
		Id int `json:"id"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("error decoding addfollow")
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	privacy, err := h.store.CheckUserPrivacyStatus(data.Id)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	fmt.Println(privacy)
	// add notification

	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	if privacy == "public" {
		fmt.Println("public user")
		exists, err := h.store.AddFollower(user.Id, data.Id, "completed")
		if err != nil {
			responseData["response"] = "failure"
			responseData["message"] = "Internal server error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
		}
		if exists == 0 {
			fmt.Println("already following")
			responseData["response"] = "failure"
			responseData["message"] = "User is already followin them"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
		}
		responseData["response"] = "success"
		responseData["following"] = "following"
		responseData["message"] = "send Follow request success"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	// user is private

	// add follower as pending
	followerTableId, err := h.store.AddFollower(user.Id, data.Id, "pending")
	if err != nil {
		fmt.Println("err getting followertableId", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	fmt.Println(followerTableId)
	if followerTableId == 0 {
		fmt.Println("already following")
		// the user is already following them
		responseData["response"] = "failure"
		responseData["message"] = "User is already followin them"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["following"] = "pending"
	responseData["message"] = "Sent a successful follow request"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting contacts")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != "GET" {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	contacts, err := h.store.GetContacts(user.Id)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "getContacts successful"
	responseData["contacts"] = contacts
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
