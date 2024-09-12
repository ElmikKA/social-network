package models

import "time"

type Session struct {
	Id      int       `json:"id"`
	Cookie  string    `json:"cookie"`
	Expires time.Time `json:"expires"`
}
