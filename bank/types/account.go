package types

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
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
