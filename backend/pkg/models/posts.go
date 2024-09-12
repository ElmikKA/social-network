package models

type Post struct {
	Id        string `json:"id"`
	UserId    int    `json:"userId"`
	GroupId   int    `json:"groupId"`
	Creator   string `json:"creator"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
	Privacy   string `json:"privacy"`
}
