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

func (h *Handler) AddPost(w http.ResponseWriter, r *http.Request) {
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

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	var user models.Users
	user, err = h.store.GetUserFromCookie(r)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	GroupId, _ := strconv.Atoi(r.FormValue("groupId"))

	post := models.Post{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
		Privacy: r.FormValue("privacy"),
		Creator: user.Name,
		UserId:  user.Id,
		GroupId: GroupId,
	}

	_, _, err = r.FormFile("avatar")
	if err == nil {
		fmt.Println("has avatar")
		filepath, err := utils.SaveFile(r, post.Creator, "Post")
		if err != nil {
			responseData["response"] = "failure"
			responseData["message"] = "Internal server error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
		}
		post.Avatar = filepath
	}

	err = h.store.AddPost(post)
	if err != nil {
		http.Error(w, "Unable to add post", http.StatusInternalServerError)
		return
	}
	responseData["response"] = "success"
	responseData["message"] = "successfully added post"
	responseData["post"] = post

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != http.MethodGet {
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

	post, err := h.store.GetPost(data.Id)

	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "GetPost success"
	responseData["post"] = post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != http.MethodGet {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	posts, err := h.store.GetAllNormalPosts()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows")
			responseData["response"] = "success"
			responseData["message"] = "GetAllPosts success"
			responseData["getAllPosts"] = nil
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
		}
		fmt.Println("error getallposts", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		responseData["getAllPosts"] = nil
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	responseData["response"] = "success"
	responseData["message"] = "GetAllPosts success"
	responseData["getAllPosts"] = posts
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)

	responseData := make(map[string]interface{})
	responseData["loggedIn"] = true
	if r.Method != http.MethodGet {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	groupData, err := h.store.GetAllGroups()
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "GetAllGroups success"
	responseData["groupData"] = groupData

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
