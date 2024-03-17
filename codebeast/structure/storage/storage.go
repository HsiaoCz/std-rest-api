package storage

import "github.com/HsiaoCz/std-rest-api/codebeast/structure/types"

type Store interface {
	GetUserByID(string) *types.User
}

