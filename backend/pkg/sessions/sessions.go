package sessions

import (
	"database/sql"
	"fmt"
	"net/http"
	"social-network/db"
	"social-network/pkg/models"
	"time"

	"github.com/google/uuid"
)

func AddSession(w http.ResponseWriter, r *http.Request, Db *sql.DB, id int) error {
	// create and add a new session
	cookie, err := r.Cookie("session")
	if err != nil || cookie == nil {
		// no cookie found or error finding cookie
		fmt.Println("no cookie found")
		cookie = &http.Cookie{
			Name:    "session",
			Value:   uuid.New().String(),
			Path:    "/",
			Expires: time.Now().Add(30 * time.Minute),
			// SameSite: http.SameSiteNoneMode,
			// Secure:   true,
		}
		http.SetCookie(w, cookie)

		session := models.Session{
			Id:      id,
			Cookie:  cookie.Value,
			Expires: cookie.Expires,
		}
		if err = db.CreateSession(Db, session); err != nil {
			fmt.Println("AddSession error creating new session", err)
			return err
		}
		fmt.Println("added new session")
	} else {
		// Cookie exiss
		session, err := db.GetSessionByCookie(Db, cookie.Value)
		if err != nil {
			// no session with that cookie in db
			fmt.Println("no session with that cookie in db")
			cookie = &http.Cookie{
				Name:    "session",
				Value:   uuid.New().String(),
				Path:    "/",
				Expires: time.Now().Add(30 * time.Minute),
				// SameSite: http.SameSiteNoneMode,
				// Secure:   true,
			}
			http.SetCookie(w, cookie)
			session = models.Session{
				Id:      id,
				Cookie:  cookie.Value,
				Expires: cookie.Expires,
			}
			if err = db.CreateSession(Db, session); err != nil {
				fmt.Println("error creating session")
				return err
			}
			fmt.Println("created new session")
		} else if session.Id != id {
			// session belongs to a different user
			if err = db.DeleteSession(Db, session.Id); err != nil {
				fmt.Println("AddSession error deleting session", err)
				return err
			}

			// create new cookie and session
			cookie = &http.Cookie{
				Name:    "session",
				Value:   uuid.New().String(),
				Path:    "/",
				Expires: time.Now().Add(30 * time.Minute),
				// SameSite: http.SameSiteNoneMode,
				// Secure:   true,
			}
			http.SetCookie(w, cookie)

			session = models.Session{
				Id:      id,
				Cookie:  cookie.Value,
				Expires: cookie.Expires,
			}
			if err = db.CreateSession(Db, session); err != nil {
				fmt.Println("error creating session")
				return err
			}
			fmt.Println("created new session")
		} else {
			// session belongs to the user
			// extends the expiresAt
			if err = db.ExtendSessionDate(Db, cookie.Value); err != nil {
				fmt.Println("AddSession error extending session")
				return err
			}
			cookie.Expires = time.Now().Add(30 * time.Minute)
			cookie.Path = "/"
			http.SetCookie(w, cookie)
			fmt.Println("extended session")
		}
	}
	return nil
}
