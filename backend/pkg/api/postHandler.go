package api

import (
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
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "unable to parse form", http.StatusBadRequest)
		return
	}

	var user models.Users
	user, err = h.store.GetUserFromCookie(r)
	if err != nil {
		http.Error(w, "addpost unable to get user", http.StatusInternalServerError)
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
			fmt.Println("error saving file", err)
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
	responseData := make(map[string]interface{})
	responseData["response"] = "success"
	responseData["post"] = post

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	url, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println("")
	}
	fmt.Println(url)
	post, err := h.store.GetPost(url)

	if err != nil {
		http.Error(w, "Invalid url path", http.StatusBadRequest)
		return
	}
	responseData := make(map[string]interface{})

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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	posts, err := h.store.GetAllPosts()
	if err != nil {
		fmt.Println("error getallposts", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	responseData := make(map[string]interface{})
	responseData["response"] = "success"
	responseData["message"] = "GetAllPosts success"
	responseData["getAllPosts"] = posts
	fmt.Println(posts[0].Creator)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
