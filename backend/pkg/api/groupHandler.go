package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
)

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating group")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var group models.Group
	json.NewDecoder(r.Body).Decode(&group)

	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	group.UserId = user.Id
	fmt.Println(group)

	// create new group

	groupId, err := h.store.AddGroup(group)
	if err != nil {
		fmt.Println("error adding group", err)
		// if already exists then respond like that
		w.WriteHeader(http.StatusConflict)
		return
	}

	group.Id = groupId

	err = h.store.AddGroupMember(group)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
