package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HsiaoCz/std-rest-api/bank/types"
	"github.com/gorilla/mux"
)

type ErrResp struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

type handlefunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(hf handlefunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hf(w, r); err != nil {
			// handle the error
			WriteJSON(w, http.StatusBadRequest, &ErrResp{
				Error: err.Error(),
				Code:  http.StatusBadRequest,
			})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/acount", makeHTTPHandleFunc(s.handleAccount))
	log.Println("JSON API server is running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
	return nil
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := types.NewAccount("anthony", "gg")
	return WriteJSON(w, http.StatusOK, account)
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
