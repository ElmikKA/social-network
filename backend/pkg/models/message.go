package models

type Message struct {
	UserId     int    `json:"userId"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	ReceiverId int    `json:"receiverId"`
	GroupId    int    `json:"groupId"`
}
