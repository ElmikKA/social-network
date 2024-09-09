package utils

import (
	"database/sql"
	"fmt"
)

func ResetOnline(Db *sql.DB) {
	query := `UPDATE users
			SET online = -1`
	_, err := Db.Exec(query)
	if err != nil {
		fmt.Println("error reseting online status", err)
	}
}

func GoOffline(Db *sql.DB, id int) error {
	query := `UPDATE users SET online = -1 WHERE id = ?`
	_, err := Db.Exec(query, id)
	if err != nil {
		fmt.Println("error changing to offline status", err)
		return err
	}
	return nil
}

func GoOnline(Db *sql.DB, id int) error {
	query := `UPDATE users SET online = 1 WHERE id = ?`
	_, err := Db.Exec(query, id)
	if err != nil {
		fmt.Println("error changing online status", err)
		return err
	}
	return nil
}
