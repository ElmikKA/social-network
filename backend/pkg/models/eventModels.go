package models

import "time"

type Event struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId"`
	GroupId     int       `json:"groupId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}
