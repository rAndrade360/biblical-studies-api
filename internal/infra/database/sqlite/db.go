package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
