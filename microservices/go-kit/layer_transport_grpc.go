package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"gokit/pb"
)

func NewGRPCServer(endpoint endpoint.Endpoint) pb.GreeterServer {
	greeter := grpc.NewServer(endpoint, decodeGRPC, encodeGRPC)
	return &grpcServer{greeter: greeter}
}

func encodeGRPC(_ context.Context, response interface{}) (interface{}, error) {
	rsp := response.(GreetingResponse)
	return &pb.GreetingResponse{Greeting: rsp.Greeting}, nil
}

func decodeGRPC(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GreetingRequest)
	return GreetingRequest{Name: req.Name}, nil
}

// implementation pb.GreeterServer
type grpcServer struct {
	greeter grpc.Handler
}

func (s *grpcServer) Greeting(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	_, res, err := s.greeter.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GreetingResponse), nil
}
