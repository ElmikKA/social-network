package models

type Users struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	AboutMe     string `json:"aboutMe"`
	Online      string `json:"online"`
	Privacy     string `json:"privacy"`
}

type LoginCredentials struct {
	Name string `json:"name"`
	Pass string `json:"password"`
}

type GroupMember struct {
	Id     int
	Name   string
	Avatar string
	Role   string
}

type GroupEvents struct {
	Id          int
	OwnerId     int
	GroupId     int
	Title       string
	Description string
	Time        string
	Status      string
	Role        string
	IdRef int
}
