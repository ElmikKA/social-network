package models

type FollowerResponse struct {
	Pending string `json:"pending"`
	UserId  int    `json:"userId"`
}
