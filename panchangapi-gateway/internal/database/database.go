package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(databaseString string) {
	var err error

	DB, err = sql.Open("postgres", databaseString)
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("[SUCCESS] DB Connected")
}
