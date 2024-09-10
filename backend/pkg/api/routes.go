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
	router.HandleFunc("/api/getUsers", h.requireLogin(h.GetUsers))
	router.HandleFunc("/api/websocket", h.requireLogin(h.Websocket))

	// goroutine for webscocket connections
	go h.HandleWebsocketConnections()

	return router
}
