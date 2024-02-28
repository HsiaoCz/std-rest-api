package handlers

import "github.com/HsiaoCz/std-rest-api/templ-htmx/store"

type Handler struct {
	store *store.Storage
}

func New(store *store.Storage) *Handler {
	return &Handler{
		store: store,
	}
}
