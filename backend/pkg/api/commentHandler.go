package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
	"strconv"
)

func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
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

	postId, err := strconv.Atoi(r.FormValue("postId"))
	if err != nil {
		http.Error(w, "unable to parse form", http.StatusBadRequest)
		return
	}
	fmt.Println(postId)
	comment := models.Comment{
		Content: r.FormValue("content"),
		Creator: user.Name,
		UserId:  user.Id,
		PostId:  postId,
	}
	fmt.Println(comment)

	_, _, err = r.FormFile("avatar")
	if err == nil {
		fmt.Println("has avatar")
		filepath, err := utils.SaveFile(r, comment.Creator, "Post")
		if err != nil {
			fmt.Println("error saving file", err)
		}
		comment.Avatar = filepath
	} else {
		fmt.Println("no avatar")
	}

	err = h.store.AddComment(comment)
	if err != nil {
		http.Error(w, "Unable to add post", http.StatusInternalServerError)
		return
	}
	responseData := make(map[string]interface{})
	responseData["response"] = "success"
	responseData["comment"] = comment

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetComments(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	postId, err := strconv.Atoi(r.PathValue("postId"))
	if err != nil {
		http.Error(w, "Unable to add post", http.StatusInternalServerError)
		return
	}

	comments, err := h.store.GetComments(postId)
	if err != nil {
		http.Error(w, "err getting comments", http.StatusInternalServerError)
		return
	}

	responseData := make(map[string]interface{})
	responseData["response"] = "success"
	responseData["message"] = "getting comments successfully"
	responseData["comments"] = comments

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
