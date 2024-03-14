package gg

import (
	"context"
	"net"

	"github.com/HsiaoCz/std-rest-api/gpf/proto"
	"google.golang.org/grpc"
)

func MakeGRPCServer(listenAddr string, svc PriceService) error {
	grpcPriceFetcher := NewGRPCPriceFetcher(svc)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)
	return server.Serve(ln)
}

type GRPCPriceFetcher struct {
	proto.UnimplementedPriceFetcherServer
	svc PriceService
}

func NewGRPCPriceFetcher(svc PriceService) *GRPCPriceFetcher {
	return &GRPCPriceFetcher{
		svc: svc,
	}
}

func (g *GRPCPriceFetcher) FetchPrice(ctx context.Context, req *proto.PriceRequest) (resp *proto.PriceResponse, err error) {
	price, err := g.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}
	resp = &proto.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}
	return resp, err
}
