package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
)

func CheckUserExists(Db *sql.DB, user models.Users) (bool, error) {
	query := `
		SELECT COUNT(*) FROM users WHERE email = ? OR name = ? 
	`
	var count int
	err := Db.QueryRow(query, user.Email, user.Name).Scan(&count)
	if err != nil {
		fmt.Println("error checking existing user", err)
		return false, err
	}
	return count > 0, nil
}

func AddUser(Db *sql.DB, user models.Users) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println("error hashing password", err)
		return err
	}
	fmt.Println("hashed pass:", hashedPassword)
	query := `
		INSERT INTO users (name, email, password, firstName, lastName, dateOfBirth, avatar, avatarMimeType, nickname, aboutMe) VALUES (?,?,?,?,?,?,?,?,?,?)
	`
	_, err = Db.Exec(query, user.Name, user.Email, hashedPassword, user.FirstName, user.LastName, user.DateOfBirth, user.Avatar, user.AvatarMimeType, user.Nickname, user.AboutMe)
	if err != nil {
		fmt.Println("error adding user", err)
		return err
	}
	return nil
}

func CheckLogin(Db *sql.DB, credentials models.LoginCredentials) (bool, int, error) {
	query := `
		SELECT id, password FROM users WHERE (name = ? OR email = ?)
	`
	var id int
	var dbPass string
	err := Db.QueryRow(query, credentials.Name, credentials.Name).Scan(&id, &dbPass)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("user not found", err)
			return false, 0, nil
		}
		fmt.Println("error checking login", err)
		return false, 0, err
	}
	err = utils.CheckPasswordHash(credentials.Pass, dbPass)
	if err != nil {
		return false, 0, nil
	}
	return true, id, nil
}

func GetUserFromCookie(r *http.Request, Db *sql.DB) (int, string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("error getting userfromcookie", err)
		return 0, "", err
	}

	session, err := GetSessionByCookie(Db, cookie.Value)
	if err != nil {
		fmt.Println("getuserfromcookie session error", err)
		return 0, "", err
	}

	user := models.Users{}
	query := `SELECT id, name FROM users WHERE id = ?`
	err = Db.QueryRow(query, session.Id).Scan(&user.Id, &user.Name)
	if err != nil {
		fmt.Println("getuserfromcookie error", err)
		return 0, "", err
	}
	return user.Id, user.Name, nil
}
