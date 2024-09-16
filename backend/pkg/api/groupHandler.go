package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"

	"github.com/mattn/go-sqlite3"
)

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating group")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != http.MethodPost {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	var group models.Group
	json.NewDecoder(r.Body).Decode(&group)

	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	group.UserId = user.Id

	groupId, err := h.store.AddGroup(group)
	if err != nil {
		responseData["response"] = "failure"

		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
			responseData["message"] = "Group already exists with that title"
		} else {
			responseData["message"] = "Internal server error"
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	group.Id = groupId

	_, err = h.store.AddGroupMember(group)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	w.WriteHeader(http.StatusOK)
	responseData["message"] = "Group created successfully"
	responseData["response"] = "success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) RequestGroupJoin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sending group join request")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != http.MethodPost {
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
		fmt.Println("error decoding getpost")
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
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

	group := models.Group{
		UserId: user.Id,
		Id:     data.Id,
	}

	groupMembersTableId, err := h.store.AddGroupMember(group)
	if err != nil {
		responseData["response"] = "failure"
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
			responseData["message"] = "Already sent a request"
		} else {
			responseData["message"] = "Internal server error"
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	if groupMembersTableId == 0 {
		responseData["response"] = "failure"
		responseData["message"] = "User already in the group"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "GroupJoinRequest successfully sent"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
