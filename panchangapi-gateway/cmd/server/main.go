package main

import (
	"fmt"
	"net/http"
	"panchangapi-gateway/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("PanchangAPI Gateway")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "PanchangAPI Gateway")
	})

	e.GET("/sendverificationemail", handlers.VerifyEmail)
	e.Logger.Fatal(e.Start("localhost:1323"))
}
