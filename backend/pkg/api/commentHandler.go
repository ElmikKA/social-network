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
	responseData := make(map[string]interface{})
	if r.Method != "POST" {
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

	postId, err := strconv.Atoi(r.FormValue("postId"))
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Invalid payload"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}
	comment := models.Comment{
		Content: r.FormValue("content"),
		Creator: user.Name,
		UserId:  user.Id,
		PostId:  postId,
	}

	_, _, err = r.FormFile("avatar")
	if err == nil {
		fmt.Println("has avatar")
		filepath, err := utils.SaveFile(r, comment.Creator, "Post")
		if err != nil {
			fmt.Println("error saving file", err)
			responseData["response"] = "failure"
			responseData["message"] = "Internal server error"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseData)
			return
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
	responseData["response"] = "success"
	responseData["message"] = "AddComment success"
	responseData["comment"] = comment
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func (h *Handler) GetComments(w http.ResponseWriter, r *http.Request) {
	CorsEnabler(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	responseData := make(map[string]interface{})
	if r.Method != http.MethodPost {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
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

	comments, err := h.store.GetComments(data.Id)
	if err != nil {
		responseData["response"] = "failure"
		responseData["message"] = "Internal server error"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
		return
	}

	responseData["response"] = "success"
	responseData["message"] = "getting comments successfully"
	responseData["comments"] = comments
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
