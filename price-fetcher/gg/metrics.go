package gg

import (
	"context"
	"fmt"
)

type MetricService struct {
	netx PriceFetcher
}

func (s *MetricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("pushing metrics to prometheus")
	// your metrics stoarge Push to promethues (gauge counters)
	return s.netx.FetchPrice(ctx, ticker)
}
