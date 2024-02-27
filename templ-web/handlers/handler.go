package handlers

import (
	"net/http"

	"github.com/HsiaoCz/std-rest-api/templ-web/store"
	"github.com/gorilla/mux"
)

type Handler struct {
	addr  string
	store store.Store
	user  *UserHandler
}

func NewHandler(addr string, store store.Store, user *UserHandler) *Handler {
	return &Handler{
		addr:  addr,
		store: store,
		user:  user,
	}
}

func (h *Handler) Serve() error {
	r := mux.NewRouter()
	user := r.PathPrefix("/api/v1/user").Subrouter()
	user.HandleFunc("/signup", h.user.handlerUserSignup).Methods("POST")
	user.HandleFunc("/login", h.user.handleUserLogin).Methods("POST")
	user.HandleFunc("/hello", h.user.handleUserSayHello).Methods("POST")
	return http.ListenAndServe(h.addr, r)
}
