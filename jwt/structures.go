package jwt

import (
	"errors"
	"strings"
)

type HeaderToken struct {
	Authorization string `json:"Authorization" from:"Authorization"`
}

func (ht *HeaderToken) GetToken() (string, error) {
	// Check if the token starts with "Bearer"
	parts := strings.Fields(ht.Authorization)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}
