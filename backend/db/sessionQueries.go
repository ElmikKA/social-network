package db

import (
	"database/sql"
	"fmt"
	"social-network/pkg/models"
	"time"
)

func CreateSession(Db *sql.DB, session models.Session) error {
	query := `
		INSERT OR REPLACE INTO sessions (userId, cookie, expiresAt) VALUES (?,?,?)
	`
	_, err := Db.Exec(query, session.Id, session.Cookie, session.Expires)
	if err != nil {
		fmt.Println("error creatin session", err)
		return err
	}
	fmt.Println("inserted or replaced a new session")
	return nil
}

func DeleteSession(Db *sql.DB, id int) error {
	query := `
		DELETE FROM sessions WHERE userId = ?
	`
	_, err := Db.Exec(query, id)
	if err != nil {
		fmt.Println("error deleting session", err)
		return err
	}
	return nil
}

func GetSessionByCookie(Db *sql.DB, cookie string) (models.Session, error) {
	query := `SELECT userId, expiresAt FROM sessions WHERE cookie = ?`
	session := models.Session{}
	err := Db.QueryRow(query, cookie).Scan(&session.Id, &session.Expires)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("getsessionbycookie session not found", err)
			return session, err
		}
		fmt.Println("getSessionByCookie error", err)
		return session, err
	}
	return session, nil
}

func ExtendSessionDate(Db *sql.DB, cookie string) error {
	query := `UPDATE sessions SET expiresAt = ? WHERE cookie = ?`

	_, err := Db.Exec(query, time.Now().Add(30*time.Minute), cookie)
	if err != nil {
		fmt.Println("extendSessionDate errr", err)
		return err
	}
	return nil
}
