package db

import (
	"database/sql"
	"fmt"
	"social-network/pkg/models"
	"time"
)

func (s *Store) CreateSession(session models.Session) error {
	query := `
		INSERT OR REPLACE INTO sessions (userId, cookie, expiresAt) VALUES (?,?,?)
	`
	_, err := s.Db.Exec(query, session.Id, session.Cookie, session.Expires)
	if err != nil {
		fmt.Println("error creatin session", err)
		return err
	}
	fmt.Println("inserted or replaced a new session")
	return nil
}

func (s *Store) DeleteSession(id int) error {
	query := `
		DELETE FROM sessions WHERE userId = ?
	`
	_, err := s.Db.Exec(query, id)
	if err != nil {
		fmt.Println("error deleting session", err)
		return err
	}
	return nil
}

func (s *Store) GetSessionByCookie(cookie string) (models.Session, error) {
	query := `SELECT userId, expiresAt FROM sessions WHERE cookie = ?`
	session := models.Session{}
	err := s.Db.QueryRow(query, cookie).Scan(&session.Id, &session.Expires)
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

func (s *Store) ExtendSessionDate(cookie string) error {
	query := `UPDATE sessions SET expiresAt = ? WHERE cookie = ?`

	_, err := s.Db.Exec(query, time.Now().Add(30*time.Minute), cookie)
	if err != nil {
		fmt.Println("extendSessionDate errr", err)
		return err
	}
	return nil
}
