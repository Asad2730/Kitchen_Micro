package main

import (
	"net"

	handler "github.com/Asad2730/Kitchen_Micro/services/orders/handler/orders"
	"github.com/Asad2730/Kitchen_Micro/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(add string) *gRPCServer {
	return &gRPCServer{addr: add}
}

func (s *gRPCServer) Run() error {

	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	gRPCServer := grpc.NewServer()

	//register our gRPC services
	orderService := service.NewOrderService()
	handler.NewGrpcHandler(gRPCServer, orderService)

	if err := gRPCServer.Serve(listen); err != nil {
		return err
	}
	return nil
}
