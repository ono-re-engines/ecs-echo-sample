package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	for _, arg := range os.Args {
		if arg == "test" {
			fmt.Printf("sleep starting at %s", time.Now().Format(time.RFC3339))
			time.Sleep(time.Minute * 5)
			fmt.Printf("sleep ended at %s", time.Now().Format(time.RFC3339))
			return
		}
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		msg := os.Getenv("MESSAGE")
		if msg == "" {
			msg = "Env not found."
		}
		return c.String(http.StatusOK, msg)
	})
	e.GET("/api/test", func(c echo.Context) error {
		msg := os.Getenv("MESSAGE")
		if msg == "" {
			msg = "Env not found."
		}
		return c.String(http.StatusOK, msg)
	})

	e.Logger.Fatal(e.Start(":80"))
}
