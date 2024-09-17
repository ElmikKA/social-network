package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() (*sql.DB, error) {
	Db, err := sql.Open("sqlite3", "./db/sn.db")
	if err != nil {
		return nil, err
	}

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

	// add triggers if they don't exist

	return Db, nil
}
