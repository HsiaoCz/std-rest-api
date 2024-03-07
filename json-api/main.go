package main

import (
	"encoding/json"
	"net/http"
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
	http.ListenAndServe(":9001", nil)
}

type User struct {
	ID       int
	Username string
	Vaild    bool
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
