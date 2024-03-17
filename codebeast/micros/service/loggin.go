package service

import (
	"context"
	"fmt"
	"time"

	"github.com/HsiaoCz/std-rest-api/codebeast/micros/types"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (l *LoggingService) GetCatFact(ctx context.Context) (fact *types.CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())
	return l.next.GetCatFact(ctx)
}
