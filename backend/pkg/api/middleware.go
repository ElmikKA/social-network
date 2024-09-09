package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/db"
)

func (api *APIServer) requireLogin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err == nil {
			// cookie found
			_, err = db.GetSessionByCookie(api.db, cookie.Value)
			if err == nil {
				// logged in
				fmt.Println("middleware: logged in")
				var err error
				api.id, api.username, err = db.GetUserFromCookie(r, api.db)
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
