package api

import (
	"net/http"
)

func CorsEnabler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	isAllowed := true
	if isAllowed {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Header.Get("Upgrade") == "websocket" {
			w.Header().Set("Connection", "Upgrade")
			w.Header().Set("Upgrade", "websocket")
		}
	} else {
		http.Error(w, "CORS not allowed", http.StatusForbidden)
	}
}
