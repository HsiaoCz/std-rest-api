package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/std-rest-api/codebeast/micros/types"
)

type Service interface {
	GetCatFact(context.Context) (*types.CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*types.CatFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	fact := &types.CatFact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}
	return fact, nil
}
