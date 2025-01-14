package api

import (
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) AddSession(w http.ResponseWriter, r *http.Request, id int) error {
	CorsEnabler(w, r)
	// create and add a new session
	cookie, err := r.Cookie("session")
	if err != nil || cookie == nil {
		// no cookie found or error finding cookie
		cookie = &http.Cookie{
			Name:     "session",
			Value:    uuid.New().String(),
			Path:     "/",
			Expires:  time.Now().Add(30 * time.Minute),
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, cookie)

		session := models.Session{
			Id:      id,
			Cookie:  cookie.Value,
			Expires: cookie.Expires,
		}
		if err = h.store.CreateSession(session); err != nil {
			fmt.Println("AddSession error creating new session", err)
			return err
		}
	} else {
		// Cookie exiss
		session, err := h.store.GetSessionByCookie(cookie.Value)
		if err != nil {
			// no session with that cookie in db
			cookie = &http.Cookie{
				Name:     "session",
				Value:    uuid.New().String(),
				Path:     "/",
				Expires:  time.Now().Add(30 * time.Minute),
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
			}
			http.SetCookie(w, cookie)
			session = models.Session{
				Id:      id,
				Cookie:  cookie.Value,
				Expires: cookie.Expires,
			}
			if err = h.store.CreateSession(session); err != nil {
				fmt.Println("error creating session")
				return err
			}
		} else if session.Id != id {
			// session belongs to a different user
			if err = h.store.DeleteSession(session.Id); err != nil {
				fmt.Println("AddSession error deleting session", err)
				return err
			}

			// create new cookie and session
			cookie = &http.Cookie{
				Name:     "session",
				Value:    uuid.New().String(),
				Path:     "/",
				Expires:  time.Now().Add(30 * time.Minute),
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
			}
			http.SetCookie(w, cookie)

			session = models.Session{
				Id:      id,
				Cookie:  cookie.Value,
				Expires: cookie.Expires,
			}
			if err = h.store.CreateSession(session); err != nil {
				fmt.Println("error creating session")
				return err
			}
		} else {
			// session belongs to the user
			// extends the expiresAt
			if err = h.store.ExtendSessionDate(cookie.Value); err != nil {
				fmt.Println("AddSession error extending session")
				return err
			}
			cookie.Expires = time.Now().Add(30 * time.Minute)
			cookie.Path = "/"
			http.SetCookie(w, cookie)
		}
	}
	return nil
}
