package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/std-rest-api/codebeast/micros/service"
)

type ApiServer struct {
	svc service.Service
}

func NewApiServer(svc service.Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (a *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", a.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

func (a *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := a.svc.GetCatFact(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, &fact)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
