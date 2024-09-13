package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg/models"
)

func (h *Handler) requireLogin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CorsEnabler(w, r)
		cookie, err := r.Cookie("session")
		if err == nil {
			// cookie found
			_, err = h.store.GetSessionByCookie(cookie.Value)
			if err == nil {
				// logged in
				fmt.Println("middleware: logged in")
				var err error
				var user models.Users
				user, err = h.store.GetUserFromCookie(r)
				h.username = user.Name
				h.id = user.Id
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
		responseData["response"] = "failure"
		responseData["message"] = "not logged in"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
	}
}
