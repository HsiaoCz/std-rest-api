package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HsiaoCz/std-rest-api/bank/store"
	"github.com/HsiaoCz/std-rest-api/bank/types"
	"github.com/gorilla/mux"
)

type ErrResp struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

type MsgResp struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Response struct {
	Data    any     `json:"data"`
	MsgResp MsgResp `json:"Msg_Resp"`
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
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
	store      store.Storage
}

func NewAPIServer(listenAddr string, store store.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))
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
	id := mux.Vars(r)["id"]
	fmt.Println("account id :", id)
	account := types.NewAccount("anthony", "gg")
	return WriteJSON(w, http.StatusOK, account)
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := types.CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		return WriteJSON(w, http.StatusBadRequest, ErrResp{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
	}
	return WriteJSON(w, http.StatusOK, &Response{
		Data: types.NewAccount(createAccountReq.FirstName, createAccountReq.LastName),
		MsgResp: MsgResp{
			Message: "create account successed!",
			Code:    http.StatusOK,
		},
	})
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
