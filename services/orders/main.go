package main

import (
	"github.com/Asad2730/Kitchen_Micro/services/orders/server"
)

func main() {
	httpServer := server.NewHttpServer(":8000")
	go httpServer.Run()

	grpcServer := server.NewGRPCServer(":9000")
	grpcServer.Run()
}
