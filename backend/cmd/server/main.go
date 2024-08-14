package main

import (
	"fmt"
	"net/http"
	"os"
	"panchangapi-gateway/internal/api/handlers"
	"panchangapi-gateway/internal/database"
	"panchangapi-gateway/internal/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("PanchangAPI Gateway")

	utils.LoadEnv()

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

	database.InitDB(conn_string)

	// init redis by loading env variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	database.InitRedis(redisHost, redisPort)

	// setting up echo and routes, handlers
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "PanchangAPI Gateway")
	})

	e.POST("/sendverificationemail", handlers.VerifyEmail)
	e.POST("/verifyotp", handlers.VerifyOtp)
	e.POST("/createProfile", handlers.CreateProfile)

	// start server
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))

}
