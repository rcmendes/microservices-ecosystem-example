package ucs

// import (
// 	"net/http"

// 	"github.com/google/uuid"
// 	"github.com/labstack/echo"
// )

// type Credentials interface {
// 	Username() string
// 	Password() string
// }

// type Token interface {
// 	AccessToken() string
// }

// type LoginService interface {
// 	login(credentials Credentials) (Token, error)
// }

// func login(c echo.Context) error {
// 	contentType := c.Request().Header.Get("Content-Type")
// 	println(contentType)

// 	credentials := new(Credentials)

// 	if err := c.Bind(credentials); err != nil {
// 		c.String(http.StatusBadRequest, "Invalid credentials.")
// 	}

// 	if credentials.Username != "rodrigo" || credentials.Password != "senha" {
// 		return c.String(http.StatusUnauthorized, "")
// 	}

// 	token := uuid.NewString()
// 	access_tokens = append(access_tokens, token)

// 	return c.JSON(http.StatusOK, &Token{
// 		token,
// 	})
// }
