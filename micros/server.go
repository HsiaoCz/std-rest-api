package micros

import "context"

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}
