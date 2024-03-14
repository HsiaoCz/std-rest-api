package gg

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type MyKey string

var prices = map[string]float64{
	"ETH": 999.99,
	"BTC": 200000.0,
	"GG":  1000000.0,
}

// PriceService is an interface that can fetch the price for any given ticker
type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

type PriceServer struct{}

func (s *PriceServer) FetchPrice(_ context.Context, ticker string) (float64, error) {
	price, ok := prices[ticker]
	if !ok {
		return 0.0, fmt.Errorf("price for ticker (%s) is not available", ticker)
	}
	return price, nil
}

type LoggingService struct {
	PriceServer
}

func (s LoggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		reqID := ctx.Value(MyKey("requestID"))
		logrus.WithFields(logrus.Fields{
			"requestID": reqID,
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())
	return s.PriceServer.FetchPrice(ctx, ticker)
}
