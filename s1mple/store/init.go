package store

import "github.com/HsiaoCz/std-rest-api/s1mple/types"

type Store interface {
	CreateUser(*types.User) error
}
