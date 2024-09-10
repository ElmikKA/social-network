package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) requireLogin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err == nil {
			// cookie found
			_, err = h.store.GetSessionByCookie(cookie.Value)
			if err == nil {
				// logged in
				fmt.Println("middleware: logged in")
				var err error
				h.id, h.username, err = h.store.GetUserFromCookie(r)
				if err != nil {
					fmt.Println("error getting id and username on middleware", err)
					return
				}
				handler.ServeHTTP(w, r)
				return
			}
		}
		fmt.Println("middleware: not logged in")
		responseData := make(map[string]interface{})
		responseData["requireLogin"] = "failure"
		responseData["message"] = "not logged in"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
	}
}
