package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HsiaoCz/std-rest-api/gpf/proto"
	"github.com/HsiaoCz/std-rest-api/gpf/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// type Client interface {
// 	FetchPrice()
// }

type Client struct {
	endpoint string
}

func New(enpoint string) *Client {
	return &Client{
		endpoint: enpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service response with none OK status code: %v", httpErr["error"])
	}
	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}
	return priceResp, nil
}

func NewGRPCClient(remoteAddr string) (proto.PriceFetcherClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := proto.NewPriceFetcherClient(conn)
	return c, nil
}
