package models

import "time"

type Users struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	Avatar         string    `json:"avatar"`
	AvatarMimeType string    `json:"avatarMimeType"`
	Nickname       string    `json:"nickname"`
	AboutMe        string    `json:"aboutMe"`
	Online         string    `json:"online"`
}

type LoginCredentials struct {
	Name string `json:"name"`
	Pass string `json:"password"`
}

type RegisterCredentials struct {
	Name string `json:"name"`
	Pass string `json:"password"`
}