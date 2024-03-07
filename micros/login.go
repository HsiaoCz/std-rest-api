package micros

import (
	"context"
	"fmt"
	"time"
)

type LoginService struct {
	next Service
}

func NewLoginService(next Service) Service {
	return &LoginService{
		next: next,
	}
}

func (l *LoginService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%v err=%v took=%v\n", fact, err, time.Since(start))
	}(time.Now())
	return l.next.GetCatFact(ctx)
}
