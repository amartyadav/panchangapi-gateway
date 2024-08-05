package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"panchangapi-gateway/internal/api/handlers"
	"panchangapi-gateway/internal/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("PanchangAPI Gateway")

	loadEnv()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "PanchangAPI Gateway")
	})

	e.GET("/sendverificationemail", handlers.VerifyEmail)

	// init db by loading env varibles
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// form conn string for db
	conn_string := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	fmt.Println("[INFO] Connection string: ", conn_string)

	database.InitDB(conn_string)

	// start server
	e.Logger.Fatal(e.Start("localhost:1323"))

}

func loadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("db.go: No .env file found, proceeding with environment variables")
		} else {
			log.Printf("db.go: Error loading .env file: %v", err)
		}
	} else {
		log.Println("db.go: .env file loaded successfully")
	}

}
