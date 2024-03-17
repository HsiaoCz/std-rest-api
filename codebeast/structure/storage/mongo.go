package storage

import "github.com/HsiaoCz/std-rest-api/codebeast/structure/types"

type MongoStorage struct{}

func NewMongoStorage() *MongoStorage {
	return &MongoStorage{}
}

func (s *MongoStorage) GetUserByID(id string) *types.User {
	return &types.User{
		ID:   "1223333",
		Name: "jesca",
	}
}
