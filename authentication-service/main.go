package main

import (
	"authentication-service/adapters/web"
	"authentication-service/ucs"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	// TODO Needs to add a mutex
	tokens := make([]string, 0)

	loginUC, err := ucs.NewLoginService(&tokens)
	if err != nil {
		log.Fatal(err)
	}

	server.POST("/login", web.Login(loginUC))

	server.Logger.Fatal(server.Start(":9001"))
}
