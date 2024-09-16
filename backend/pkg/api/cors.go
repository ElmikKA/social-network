package api

import (
	"net/http"
)

func CorsEnabler(w http.ResponseWriter, r *http.Request) {
	// allowedOrigins := []string{
	// 	"http://localhost:8080",
	// 	"http://localhost:5173",
	// }
	origin := r.Header.Get("Origin")
	// isAllowed := false
	// for _, o := range allowedOrigins {
	// 	if origin == o {
	// 		isAllowed = true
	// 		break
	// 	}
	// }

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
	} else {
		// Optionally handle disallowed origins (e.g., returning an error)
		http.Error(w, "CORS not allowed", http.StatusForbidden)
	}
}
