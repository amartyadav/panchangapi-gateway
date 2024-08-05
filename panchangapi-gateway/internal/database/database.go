package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(databaseString string) {
	var err error

	db, err := sql.Open("postgres", databaseString)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("[SUCCESS] DB Connected")
}
