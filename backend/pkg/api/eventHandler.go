package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating event")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	if r.Method != http.MethodPost {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)
	fmt.Println(event)

	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	event.UserId = user.Id

	err = h.store.AddEvent(event)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	// add trigger for adding owner to members

	responseData["response"] = "success"
	responseData["message"] = "successfully created event"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)

}
