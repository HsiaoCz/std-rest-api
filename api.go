package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	r := mux.NewRouter()
	subRouter := r.PathPrefix("/api/v1").Subrouter()

	zap.L().Info("Starting the API server", zap.String("addr", s.addr))
	http.ListenAndServe(s.addr, subRouter)
}
