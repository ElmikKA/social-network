package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() (*sql.DB, error) {
	Db, err := sql.Open("sqlite3", "./db/sn.db")
	if err != nil {
		return nil, err
	}

	// add triggers if they don't exist

	return Db, nil
}
