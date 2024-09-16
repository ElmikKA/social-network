package models

type FollowerResponse struct {
	Pending string `json:"pending"`
	UserId  int    `json:"userId"`
}

type Contacts struct {
	Id     int
	Name   string
	Avatar string
	Type   string
}

type GroupChats struct {
	GroupId int
	Title   string
}
