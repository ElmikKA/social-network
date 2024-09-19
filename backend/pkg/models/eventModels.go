package models

type Event struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	GroupId     int    `json:"groupId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Time        string `json:"time"`
}

type EventResponse struct {
	EventId int    `json:"eventId"`
	UserId  int    `json:"userId"`
	Pending string `json:"pending"`
}
