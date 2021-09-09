package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoMessage struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.GET("/:message", func(c echo.Context) error {
		message := &EchoMessage{
			Message: c.Param("message"),
		}
		return c.JSON(http.StatusOK, message)
	})
	e.Logger.Fatal(e.Start(":9000"))
}
