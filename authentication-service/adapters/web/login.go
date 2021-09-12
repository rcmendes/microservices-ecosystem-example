package web

import (
	"authentication-service/ucs"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Credentials struct {
	User string `json:"username"`
	Pwd  string `json:"password"`
}

func (c *Credentials) Username() string {
	return c.User
}

func (c *Credentials) Password() string {
	return c.Pwd
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type EchoFunction = func(c echo.Context) error

func Login(loginUC ucs.LoginService) EchoFunction {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")
		println(contentType)

		credentials := new(Credentials)

		if err := c.Bind(credentials); err != nil {
			c.String(http.StatusBadRequest, "Invalid credentials.")
		}

		token, err := loginUC.Execute(credentials)
		if err != nil {
			if errors.Is(err, ucs.ErrInvalidCredentials) {
				code := http.StatusUnauthorized
				message := LoginError(code, err)
				return c.JSON(code, message)
			}
		}

		return c.JSON(http.StatusOK, &Token{
			token.AccessToken(),
		})
	}
}
