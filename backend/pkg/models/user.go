package models

type Users struct {
	Id        int    `json:"id"`
	Name      string `json:"username"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Online    string `json:"online"`
}
