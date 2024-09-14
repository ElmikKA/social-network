package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) AddFollow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adding follower")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	if r.Method != "GET" {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	follow, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Invalid url payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	privacy, err := h.store.CheckUserPrivacyStatus(follow)
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
		_, err := h.store.AddFollower(user.Id, follow, "completed")
		if err != nil {
			responseData["response"] = "failure"
			responseData["message"] = "Internal server error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
		}
		responseData["response"] = "success"
		responseData["message"] = "send Follow request success"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	// user is private

	// add follower as pending
	followerTableId, err := h.store.AddFollower(user.Id, follow, "pending")
	if err != nil {
		fmt.Println("err getting followertableId", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	if followerTableId == 0 {
		// the user is already following them
		responseData["response"] = "failure"
		responseData["message"] = "User is already followin them"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	responseData["response"] = "success"
	responseData["message"] = "Sent a successful follow request"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

// func (h *Handler) RespondFollow(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("responding to follow")
// 	CorsEnabler(w, r)
// 	if r.Method == http.MethodOptions {
// 		return
// 	}
// 	responseData := make(map[string]interface{})
// 	if r.Method != "POST" {
// 		responseData["response"] = "failure"
// 		responseData["message"] = "Method not allowed"
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(responseData)
// 		return
// 	}

// 	users, err := h.store.GetUserFromCookie(r)
// 	if err != nil {
// 		responseData["response"] = "failure"
// 		responseData["message"] = "Internal server error"
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(responseData)
// 		return
// 	}

// 	var data models.FollowerResponse
// 	err = json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		fmt.Println("err decoding json respondfollow", err)
// 		responseData["response"] = "failure"
// 		responseData["message"] = "Invalid payload"
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(responseData)
// 		return
// 	}
// 	fmt.Println(data)

// 	err = h.store.RespondFollow(users.Id, data.UserId, data.Pending)
// 	if err != nil {
// 		fmt.Println("err responding to follow")
// 		responseData["response"] = "failure"
// 		responseData["message"] = "Internal server error"
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(responseData)
// 		return
// 	}
// 	responseData["response"] = "success"
// 	responseData["message"] = "respondFollow success"
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(responseData)
// }
