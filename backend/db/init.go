package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb(trigger bool) (*sql.DB, error) {
	Db, err := sql.Open("sqlite3", "./db/sn.db")
	if err != nil {
		return nil, err
	}
	if trigger {
		triggers, err := os.ReadFile("./db/triggers.sql")
		if err != nil {
			fmt.Println("error adding triggers")
			return nil, err
		}
		_, err = Db.Exec(string(triggers))
		if err != nil {
			fmt.Println("error adding triggers")
			return nil, err
		}
	}

	return Db, nil
}
