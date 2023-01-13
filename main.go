package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		msg := os.Getenv("MESSAGE")
		if msg == "" {
			msg = "Env not found."
		}
		return c.String(http.StatusOK, msg)
	})

	e.Logger.Fatal(e.Start(":80"))
}
