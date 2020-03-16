package handlers

import (
	"context"

	pb "gk-truss"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.EchoServer {
	return echoService{}
}

type echoService struct{}

// Echo implements Service.
func (s echoService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp = pb.EchoResponse{
		// Out:
	}
	return &resp, nil
}

// Louder implements Service.
func (s echoService) Louder(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp = pb.EchoResponse{
		// Out:
	}
	return &resp, nil
}

// LouderGet implements Service.
func (s echoService) LouderGet(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp = pb.EchoResponse{
		// Out:
	}
	return &resp, nil
}
