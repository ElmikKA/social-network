package models

import (
	"net/http"
)

// add every function that uses db here
type UserStore interface {
	CheckUserExists(user Users) (bool, error)
	AddUser(user Users) error
	CheckLogin(credentials LoginCredentials) (bool, int, error)
	GetUserFromCookie(r *http.Request) (int, string, error)
	GetSessionByCookie(cookie string) (Session, error)
	CreateSession(session Session) error
	DeleteSession(id int) error
	ExtendSessionDate(cookie string) error
	ResetOnline()
	GoOffline(id int) error
	GoOnline(id int) error
}
