package adapters

import (
	"authentication-service/core"

	"github.com/google/uuid"
)

type Token interface {
	AccessToken() string
}

type token struct {
	access_token string
}

func (t *token) AccessToken() string {
	return t.access_token
}

func CreateToken(credentials core.BasicCredentials) (Token, error) {
	value := uuid.NewString()

	return &token{
		value,
	}, nil
}
