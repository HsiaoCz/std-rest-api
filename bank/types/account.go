package types

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000),
		Balance:   0,
	}
}
