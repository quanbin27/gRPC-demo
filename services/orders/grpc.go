package main

import (
	"github.com/quanbin27/gRPC-demo/services/orders/handler"
	"github.com/quanbin27/gRPC-demo/services/orders/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)
	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
