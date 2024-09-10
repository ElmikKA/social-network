package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
)

func (s *Store) CheckUserExists(user models.Users) (bool, error) {
	query := `
		SELECT COUNT(*) FROM users WHERE email = ? OR name = ? 
	`
	var count int
	err := s.Db.QueryRow(query, user.Email, user.Name).Scan(&count)
	if err != nil {
		fmt.Println("error checking existing user", err)
		return false, err
	}
	return count > 0, nil
}

func (s *Store) AddUser(user models.Users) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println("error hashing password", err)
		return err
	}
	fmt.Println("hashed pass:", hashedPassword)
	query := `
		INSERT INTO users (name, email, password, firstName, lastName, dateOfBirth, avatar, avatarMimeType, nickname, aboutMe) VALUES (?,?,?,?,?,?,?,?,?,?)
	`
	_, err = s.Db.Exec(query, user.Name, user.Email, hashedPassword, user.FirstName, user.LastName, user.DateOfBirth, user.Avatar, user.AvatarMimeType, user.Nickname, user.AboutMe)
	if err != nil {
		fmt.Println("error adding user", err)
		return err
	}
	return nil
}

func (s *Store) CheckLogin(credentials models.LoginCredentials) (bool, int, error) {
	query := `
		SELECT id, password FROM users WHERE (name = ? OR email = ?)
	`
	var id int
	var dbPass string
	err := s.Db.QueryRow(query, credentials.Name, credentials.Name).Scan(&id, &dbPass)
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

func (s *Store) GetUserFromCookie(r *http.Request) (int, string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("error getting userfromcookie", err)
		return 0, "", err
	}

	session, err := s.GetSessionByCookie(cookie.Value)
	if err != nil {
		fmt.Println("getuserfromcookie session error", err)
		return 0, "", err
	}

	user := models.Users{}
	query := `SELECT id, name FROM users WHERE id = ?`
	err = s.Db.QueryRow(query, session.Id).Scan(&user.Id, &user.Name)
	if err != nil {
		fmt.Println("getuserfromcookie error", err)
		return 0, "", err
	}
	return user.Id, user.Name, nil
}
