package api

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/std-rest-api/codebeast/structure/storage"
)

type Server struct {
	listenAddr string
	store      storage.Store
}

func NewServer(listenAddr string, store storage.Store) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.GetUserByID("1222")
	json.NewEncoder(w).Encode(&user)
}
