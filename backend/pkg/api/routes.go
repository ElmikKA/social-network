package api

import (
	"net/http"
	"social-network/pkg/models"
)

type Handler struct {
	store    models.UserStore
	username string
	id       int
}

func NewHandler(store models.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Routes() *http.ServeMux {
	h.store.ResetOnline()
	router := http.NewServeMux()

	router.HandleFunc("/api/register", h.Register)
	router.HandleFunc("/api/login", h.Login)
	router.HandleFunc("/api/getUser/{userId}", h.requireLogin(h.GetUser))
	router.HandleFunc("/api/getAllUsers", h.requireLogin(h.GetAllUsers))
	router.HandleFunc("/api/avatars/", h.ServeAvatar)
	router.HandleFunc("/api/addPost", h.requireLogin(h.AddPost))
	router.HandleFunc("/api/getPost/{id}", h.requireLogin(h.GetPost))
	router.HandleFunc("/api/getAllPosts", h.requireLogin(h.GetAllPosts))
	router.HandleFunc("/api/addComment", h.requireLogin(h.AddComment))
	router.HandleFunc("/api/getComments/{postId}", h.requireLogin(h.GetComments))
	router.HandleFunc("/api/addFollow/{userId}", h.requireLogin(h.AddFollow))
	router.HandleFunc("/api/respondFollow/{userId}", h.requireLogin(h.RespondFollow))
	router.HandleFunc("/api/createGroup", h.requireLogin(h.CreateGroup))

	router.HandleFunc("/api/websocket", h.requireLogin(h.Websocket))
	// goroutine for webscocket connections
	go h.HandleWebsocketConnections()

	return router
}
