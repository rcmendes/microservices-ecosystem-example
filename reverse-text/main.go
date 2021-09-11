package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoMessage struct {
	Message string `json:"message"`
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	e := echo.New()
	e.GET("/:message", func(c echo.Context) error {
		message := &EchoMessage{
			Message: reverse(c.Param("message")),
		}
		return c.JSON(http.StatusOK, message)
	})

	e.GET("/", func(c echo.Context) error {
		message := &EchoMessage{
			Message: "Welcome to Reverse Text API!",
		}
		return c.JSON(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
