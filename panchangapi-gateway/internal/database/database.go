package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(databaseString string) error {
	var err error

	db, err := sql.Open("postgres", databaseString)
	if err != nil {
		return err
	}

	return db.Ping()
}