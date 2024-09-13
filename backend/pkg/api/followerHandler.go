package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"strconv"
)

func (h *Handler) AddFollow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adding follower")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	follow, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		http.Error(w, "Unable to add post", http.StatusInternalServerError)
		return
	}

	privacy, err := h.store.CheckUserPrivacyStatus(follow)
	if err != nil {
		fmt.Println("err checking user privacy")
		return
	}
	fmt.Println(privacy)
	// add notification

	user, err := h.store.GetUserFromCookie(r)
	if err != nil {
		fmt.Println("error getting user from cookie followers", err)
		return
	}

	if privacy == "public" {
		fmt.Println("public user")
		_, err := h.store.AddFollower(user.Id, follow, "completed")
		if err != nil {
			fmt.Println("err adding public follower", err)
			return
		}
		return
	}
	// user is private
	// needs to send notification to accept/decline follow request

	// add follower as pending and get back te id
	followerTableId, err := h.store.AddFollower(user.Id, follow, "pending")
	if err != nil {
		fmt.Println("err getting followertableId", err)
		return
	}
	if followerTableId == 0 {
		// the user is already following them
		return
	}

	content := strconv.Itoa(user.Id) + " has sent a follow request"
	notification := models.Notification{
		UserId:  follow,
		Content: content,
		Type:    "f_req",
		IdRef:   followerTableId,
	}

	err = h.store.AddNotification(notification)
	if err != nil {
		fmt.Println("error adding notification", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) RespondFollow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("responding to follow")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	responseId, err := strconv.Atoi(r.PathValue("userId"))
	if err != nil {
		http.Error(w, "invalid url type", http.StatusInternalServerError)
		return
	}

	users, err := h.store.GetUserFromCookie(r)
	if err != nil {
		http.Error(w, "unable to get user", http.StatusInternalServerError)
		return
	}

	var data models.FollowerResponse
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("err decoding followerResponse", err)
		return
	}

	err = h.store.RespondFollow(users.Id, responseId, data.Pending)
	if err != nil {
		fmt.Println("err responding to follow")
		return
	}
	fmt.Println("finished responding")
	w.WriteHeader(http.StatusOK)
}
