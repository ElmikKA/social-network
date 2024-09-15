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
	fmt.Println("adding post")
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
	fmt.Println(post)

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
	} else {
		fmt.Println("no avatar")
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
	if r.Method != http.MethodGet {
		responseData["response"] = "failure"
		responseData["message"] = "Method not allowed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	url, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println("invalid url")
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	fmt.Println(url)
	post, err := h.store.GetPost(url)

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
	fmt.Println("getallpost")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
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

func (h *Handler) GetGroupData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getallpost")
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	if r.Method != http.MethodGet {
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
		fmt.Println("User isn't part of the group", err)
		responseData["response"] = "failure"
		responseData["message"] = "User isn't part of the group"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

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

	// owner, all member, all events

	groupMembers, err := h.store.GetGroupMembers(response.GroupId)
	if err != nil {
		fmt.Println("error getting groupmembers", err)
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	events, err := h.store.GetGroupEvents(response.GroupId, h.id)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
