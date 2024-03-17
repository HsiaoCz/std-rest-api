package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHander(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("error", "err", err)
			return
		}
	}
}

type respErr struct {
	err    string
	status int
}

func (r respErr) Error() string {
	return r.err
}

func main() {
	http.HandleFunc("/user", makeHTTPHander(handleGetUserByID))
	http.ListenAndServe(":9001", nil)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return writeJSON(w, http.StatusMethodNotAllowed, respErr{err: "method not allowed!", status: http.StatusMethodNotAllowed})
	}
	return writeJSON(w, http.StatusOK, "ok")
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
