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
	if r.Method != http.MethodGet {
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

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("error decoding getpost")
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	messages, err := h.store.GetMessages(data.UserId, data.GroupId)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "getting comments successfully"
	responseData["messages"] = messages
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
