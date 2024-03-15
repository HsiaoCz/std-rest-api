package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/user", makeHTTPHandler(handleUser))
	http.ListenAndServe(":3001", nil)
}

type apifunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandler(fn apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
