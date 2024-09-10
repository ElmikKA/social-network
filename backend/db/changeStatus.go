package db

import (
	"fmt"
)

func (s *Store) ResetOnline() {
	query := `UPDATE users
			SET online = -1`
	_, err := s.Db.Exec(query)
	if err != nil {
		fmt.Println("error reseting online status", err)
	}
}

func (s *Store) GoOffline(id int) error {
	query := `UPDATE users SET online = -1 WHERE id = ?`
	_, err := s.Db.Exec(query, id)
	if err != nil {
		fmt.Println("error changing to offline status", err)
		return err
	}
	return nil
}

func (s *Store) GoOnline(id int) error {
	query := `UPDATE users SET online = 1 WHERE id = ?`
	_, err := s.Db.Exec(query, id)
	if err != nil {
		fmt.Println("error changing online status", err)
		return err
	}
	return nil
}
