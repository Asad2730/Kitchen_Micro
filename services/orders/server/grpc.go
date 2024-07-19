package server

import (
	"log"
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

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen to %v", err)
	}
	gRPCServer := grpc.NewServer()

	//register our gRPC services
	orderService := service.NewOrderService()
	handler.NewGrpcHandler(gRPCServer, orderService)

	log.Println("Started gRPC server on", s.addr)
	return gRPCServer.Serve(lis)
}
