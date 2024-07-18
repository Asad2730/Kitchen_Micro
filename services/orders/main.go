package main

import "log"

func main() {
	gRPCServer := NewGRPCServer(":9000")
	if err := gRPCServer.Run(); err != nil {
		log.Fatalf("Failed to listen to grpc server %v", err.Error())
	}
	log.Panicln("Started gRPC server on ", 9000)
}
