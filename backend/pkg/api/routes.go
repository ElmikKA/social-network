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
	router.HandleFunc("/api/avatars/", h.ServeAvatar)
	router.HandleFunc("/api/logout", h.requireLogin(h.LogOut))
	router.HandleFunc("/api/getUser/{userId}", h.requireLogin(h.GetUser))
	router.HandleFunc("/api/getAllUsers", h.requireLogin(h.GetAllUsers))
	router.HandleFunc("/api/addPost", h.requireLogin(h.AddPost))
	router.HandleFunc("/api/getPost", h.requireLogin(h.GetPost))
	router.HandleFunc("/api/getAllPosts", h.requireLogin(h.GetAllPosts))
	router.HandleFunc("/api/getGroupData", h.requireLogin(h.GetGroupData))
	router.HandleFunc("/api/getAllGroups", h.requireLogin(h.GetAllGroups))
	router.HandleFunc("/api/addComment", h.requireLogin(h.AddComment))
	router.HandleFunc("/api/getComments", h.requireLogin(h.GetComments))
	router.HandleFunc("/api/addFollow", h.requireLogin(h.AddFollow))
	router.HandleFunc("/api/createGroup", h.requireLogin(h.CreateGroup))
	router.HandleFunc("/api/requestGroupJoin", h.requireLogin(h.RequestGroupJoin))
	router.HandleFunc("/api/createEvent", h.requireLogin(h.CreateEvent))
	router.HandleFunc("/api/respondNotification", h.requireLogin(h.RespondNotification))
	router.HandleFunc("/api/getNotifications", h.requireLogin(h.GetNotifications))
	router.HandleFunc("/api/getContacts", h.requireLogin(h.GetContacts))
	router.HandleFunc("/api/websocket", h.requireLogin(h.Websocket))
	router.HandleFunc("/api/getMessages", h.requireLogin(h.GetMessages))
	router.HandleFunc("/api/checkLogin", h.requireLogin(h.CheckLogin))

	// goroutine for webscocket connections
	go h.HandleWebsocketConnections()

	return router
}
