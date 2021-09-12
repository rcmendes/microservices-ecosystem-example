package ucs

import (
	"authentication-service/adapters"
	"authentication-service/core"
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type LoginCommand interface {
	Username() string
	Password() string
}

type LoginService interface {
	Execute(command LoginCommand) (adapters.Token, error)
}

type loginService struct {
	tokenStorage *[]string
}

func NewLoginService(tokenStorage *[]string) (LoginService, error) {
	return &loginService{
		tokenStorage,
	}, nil
}

func (srv *loginService) Execute(command LoginCommand) (adapters.Token, error) {
	if command.Username() != "rodrigo" || command.Password() != "senha" {
		return nil, ErrInvalidCredentials
	}

	credentials := core.BasicCredentials{
		Username: command.Username(),
		Password: command.Password(),
	}
	token, err := adapters.CreateToken(credentials)
	if err != nil {
		return nil, err
	}

	*(srv.tokenStorage) = append(*srv.tokenStorage, token.AccessToken())

	return token, nil
}
