package api

import (
	"net/http"
)

func (api *APIServer) Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/api/register", api.Register)
	router.HandleFunc("/api/login", api.Login)
	router.HandleFunc("/api/getUsers", api.requireLogin(api.GetUsers))
	router.HandleFunc("/api/websocket", api.requireLogin(api.Websocket))

	// goroutine for webscocket connections
	go api.HandleWebsocketConnections()

	return router
}
