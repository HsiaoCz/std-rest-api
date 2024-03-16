package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type apiError struct {
	Err    string
	Status int
}

func (e apiError) Error() string {
	return e.Err
}

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHander(handler Handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJSON(w, e.Status, apiError{Err: "internal error", Status: e.Status})
			}
		}
	}
}

func main() {
	http.HandleFunc("/user", makeHTTPHander(handleGetUserById))
	http.HandleFunc("/user/create",makeHTTPHander(handleCreateUser))
	http.ListenAndServe(":9001", nil)
}

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Vaild    bool   `json:"valid"`
}

func handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}
	}
	user := User{}
	if !user.Vaild {
		return apiError{Err: "user not valid", Status: http.StatusForbidden}
	}
	return writeJSON(w, http.StatusOK, User{ID: 1, Username: "job"})
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}
	}
	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return apiError{Err: "invalid post data", Status: http.StatusBadRequest}
	}
	cuser, err := CreaterUser(user.Username, user.Vaild)
	if err != nil {
		return apiError{Err: "create user error", Status: http.StatusOK}
	}
	return writeJSON(w, http.StatusOK, cuser)
}

func CreaterUser(username string, valid bool) (*User, error) {
	return &User{
		ID:       rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100000),
		Username: username,
		Vaild:    valid,
	}, nil
}
