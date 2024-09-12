package models

type Comment struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	PostId    int    `json:"postId"`
	Creator   string `json:"creator"`
	Content   string `json:"content"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}
