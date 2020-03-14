// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	endpoint "gk-kit/hello/pkg/endpoint"
	pb "gk-kit/hello/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	foo grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.HelloServer {
	return &grpcServer{foo: makeFooHandler(endpoints, options["Foo"])}
}
