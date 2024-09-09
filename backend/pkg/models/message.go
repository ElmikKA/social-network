package models

type Message struct {
	Message    string `json:"message"`
	ReceiverId int    `json:"receiver"`
}
