package mainflux

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

func CreateUser(email, pass string) (string, error) {
	reqBody, err := json.Marshal(map[string]string{
		"email": email, "password": pass,
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost/users", "application/json",
		bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}

func GetToken(email, pass string) (Token, error) {
	reqBody, err := json.Marshal(map[string]string{
		"email": email, "password": pass,
	})
	if err != nil {
		return Token{}, err
	}

	resp, err := http.Post("http://localhost/tokens", "application/json",
		bytes.NewBuffer(reqBody))
	if err != nil {
		return Token{}, err
	}

	var t Token
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return Token{}, err
	}
	return t, nil
}
