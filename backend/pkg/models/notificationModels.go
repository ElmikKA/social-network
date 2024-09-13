package models

type Notification struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	IdRef     int    `json:"idRef"`
	CreatedAt string `json:"createdAt"`
}
