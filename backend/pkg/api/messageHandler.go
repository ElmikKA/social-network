package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {

	CorsEnabler(w, r)
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != http.MethodPost {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		fmt.Println("err getting msg user", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	var data struct {
		UserId  int `json:"userId"`
		GroupId int `json:"groupId"`
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("error decoding getMessage")
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	messages, err := h.store.GetMessages(data.UserId, data.GroupId, user.Id)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "getting messages successfully"
	responseData["messages"] = messages
	responseData["myId"] = h.id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
