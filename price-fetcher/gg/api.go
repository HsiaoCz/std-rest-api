package gg

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/HsiaoCz/std-rest-api/price-fetcher/types"
)

type APIFunction func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func makeHTTPHandler(apiFunc APIFunction) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, MyKey("requestID"), rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
				"code":  http.StatusBadRequest,
			})
		}
	}
}

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandler(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}
	priceResponse := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}
	return writeJSON(w, http.StatusOK, &priceResponse)
}

func writeJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}
