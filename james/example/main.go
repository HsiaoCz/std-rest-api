package main

import (
	"encoding/json"
	"net/http"
)

type Handler func(*Context) error

type Context struct {
	r *http.Request
	// w http.ResponseWriter
}

func POST(router string, handler Handler) {
	handler(&Context{})
}

type User struct {
	ID                 string
	Email              string
	FirstName          string
	LastName           string
	Posts              []any
	IsAdmin            bool
	EncryptedPasswrord string
	VerificationCode   int
}

type CreateUserParams struct {
	Email     string
	Passowrd  string
	FirstName string
	LastName  string
}

func handleCreateUser(c *Context) error {
	var params CreateUserParams
	if err := json.NewDecoder(c.r.Body).Decode(&params); err != nil {
		return err
	}
	var user any
	return JSON(http.StatusOK, &user)
}

func JSON(code int, v any) error {
	return nil
}

func main() {
	POST("/user", handleCreateUser)
}
