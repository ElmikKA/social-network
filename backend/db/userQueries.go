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
	query := `
		INSERT INTO users (name, email, password, firstName, lastName, dateOfBirth, avatar, nickname, aboutMe, privacy) VALUES (?,?,?,?,?,?,?,?,?,?)
	`
	_, err = s.Db.Exec(query, user.Name, user.Email, hashedPassword, user.FirstName, user.LastName, user.DateOfBirth, user.Avatar, user.Nickname, user.AboutMe, user.Privacy)
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

func (s *Store) GetUserFromCookie(r *http.Request) (models.Users, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("error getting userfromcookie", err)
		return models.Users{}, err
	}

	session, err := s.GetSessionByCookie(cookie.Value)
	if err != nil {
		fmt.Println("getuserfromcookie session error", err)
		return models.Users{}, err
	}

	user := models.Users{}
	query := `SELECT id, name, email, firstName, lastName, dateOfBirth, avatar, 
	 nickname, aboutMe, online, privacy FROM users WHERE id = ?`
	err = s.Db.QueryRow(query, session.Id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.Avatar,
		&user.Nickname,
		&user.AboutMe,
		&user.Online,
		&user.Privacy,
	)
	if err != nil {
		fmt.Println("getuserfromcookie error", err)
		return models.Users{}, err
	}
	return user, nil
}

func (s *Store) GetUser(userId int) (models.Users, error) {
	query := `SELECT id, name, email, firstName, lastName, dateOfBirth, avatar, 
	 nickname, aboutMe, online, privacy FROM users WHERE id = ?`
	var user models.Users
	err := s.Db.QueryRow(query, userId).Scan(&user.Id,
		&user.Name,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.Avatar,
		&user.Nickname,
		&user.AboutMe,
		&user.Online,
		&user.Privacy,
	)
	if err != nil {
		fmt.Println("error getting user", err)
		return models.Users{}, err
	}
	return user, nil
}

func (s *Store) GetAllUsers() ([]models.Users, error) {
	query := `
		SELECT id, name, email, firstName, lastName, dateOfBirth, avatar, 
	 nickname, aboutMe, online, privacy FROM users
	`
	rows, err := s.Db.Query(query)
	if err != nil {
		fmt.Println("error querying getallusers", err)
		return []models.Users{}, nil
	}
	defer rows.Close()

	var users []models.Users
	for rows.Next() {
		var user models.Users
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.DateOfBirth,
			&user.Avatar,
			&user.Nickname,
			&user.AboutMe,
			&user.Online,
			&user.Privacy,
		)
		if err != nil {
			fmt.Println("err scanning getallusers", err)
			return []models.Users{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *Store) CheckUserPrivacyStatus(userId int) (string, error) {
	fmt.Println(userId)
	query := `SELECT privacy FROM users WHERE id = ?`
	var privacy string
	err := s.Db.QueryRow(query, userId).Scan(&privacy)
	if err != nil {
		fmt.Println("error checking privacy status", err)
		return "", err
	}
	return privacy, nil
}

func (s *Store) ChangePrivacy(userId int, privacy string) error {
	query := `UPDATE users SET privacy =? WHERE id = ?`
	_, err := s.Db.Exec(query, privacy, userId)
	if err != nil {
		fmt.Println("err changing privacy")
		return err
	}
	return nil
}
