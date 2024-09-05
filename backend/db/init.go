package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDb() error {
	var err error
	Db, err = sql.Open("sqlite3", "./db/sn.db")
	if err != nil {
		return err
	}

	tables, err := os.ReadFile("./db/tables.sql")
	if err != nil {
		return err
	}

	_, err = Db.Exec(string(tables))
	if err != nil {
		return err
	}

	// resetOnline() reset online status whenever server closes

	return nil
}
