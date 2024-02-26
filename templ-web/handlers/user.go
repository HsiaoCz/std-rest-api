package handlers

import (
	"net/http"

	"github.com/HsiaoCz/std-rest-api/templ-web/store"
)

type UserHandler struct {
	store store.Store
}

func NewUserHandler(store store.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) handlerUserSignup(w http.ResponseWriter, r *http.Request) {}
func (u *UserHandler) handleUserLogin(w http.ResponseWriter, r *http.Request)   {}
