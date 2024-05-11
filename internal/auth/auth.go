package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKEY(header http.Header) (string, error) {
	auth := header.Get("Authorization")

	if auth == "" {
		return "", errors.New("authorization not found in header")
	}

	token := strings.Split(auth, " ")
	if len(token) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	if token[0] != "Token" {
		return "", errors.New("authorization not specific to Users")
	}

	return token[1], nil
}
