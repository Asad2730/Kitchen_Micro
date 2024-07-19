package server

import (
	"log"
	"net/http"

	handler "github.com/Asad2730/Kitchen_Micro/services/orders/handler/orders"
	"github.com/Asad2730/Kitchen_Micro/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	orderService := service.NewOrderService()
	orderhandler := handler.NewHttpHandler(orderService)
	orderhandler.RegisterRouter(router)
	log.Panicln("Starting on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
