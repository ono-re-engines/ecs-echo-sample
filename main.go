package main

import (
	"bytes"
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

	e.Logger.Fatal(e.Start(":80"))
}

func sendMessage(msg string) error {
	jsonStr := fmt.Sprintf(`{"text":"%s"}`, msg)
	req, err := http.NewRequest(
		http.MethodPost,
		"https://hooks.slack.com/services/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",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_ = resp.Body.Close()

	return nil
}
