package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"

	"github.com/mattn/go-sqlite3"
)

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
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
	responseData["groupId"] = groupId
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) RequestGroupJoin(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetGroupData(w http.ResponseWriter, r *http.Request) {
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

	var response struct {
		GroupId int `json:"groupId"`
	}
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		fmt.Println("error decoding group post", err)
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	// check if part of group

	partOfGroup, err := h.store.GetIsPartOfGroup(response.GroupId, h.id)

	if err != nil {
		fmt.Println("err getting ispartofgroup", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	if !partOfGroup {
		responseData["member"] = false
	} else {
		responseData["member"] = true
	}

	owner, err := h.store.IsGroupOwner(h.id, response.GroupId)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	responseData["owner"] = owner

	posts, err := h.store.GetAllGroupPosts(response.GroupId)
	if err != nil {
		if err == sql.ErrNoRows {
			responseData["groupPosts"] = nil
		} else {

			fmt.Println("error getallposts", err)
			responseData["response"] = "failure"
			responseData["message"] = "Internal server error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
		}
	}

	groupMembers, err := h.store.GetGroupMembers(response.GroupId)
	if err != nil {
		fmt.Println("error getting groupmembers", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	pending, err := h.store.GetGroupJoinStatus(response.GroupId, h.id)
	if err != nil {
		fmt.Println("err getting group join status", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	responseData["joinStatus"] = pending

	events, err := h.store.GetGroupEvents(response.GroupId, h.id)
	if err != nil {
		fmt.Println("error getting group events", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	group, err := h.store.GetGroup(response.GroupId)
	if err != nil {
		fmt.Println("error getting group events", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "GetAllPosts success"
	responseData["groupPosts"] = posts
	responseData["groupMembers"] = groupMembers
	responseData["groupEvents"] = events
	responseData["groupData"] = group

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetGrouInviteUsers(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)

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
		GroupId int `json:"groupId"`
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

	users, err := h.store.GetInvite(data.GroupId)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "GetAllPosts success"
	responseData["users"] = users

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) SendGroupInvite(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sending invite")
	CorsEnabler(w, r)

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
		GroupId int `json:"groupId"`
		UserId  int `json:"userId"`
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
	fmt.Println(data)

	err = h.store.SendGroupInvite(data.GroupId, data.UserId)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "Invite sent successfully"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)

}
