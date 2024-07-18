package handler

import (
	"context"

	"github.com/Asad2730/Kitchen_Micro/generated/orders"
	"github.com/Asad2730/Kitchen_Micro/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orders.UnimplementedOrderServiceServer
	orderService types.OrderService
}

func NewGrpcHandler(gRpc *grpc.Server, oderService types.OrderService) {
	ordersGrpcHandler := &OrdersGrpcHandler{orderService: oderService}

	orders.RegisterOrderServiceServer(gRpc, ordersGrpcHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{Status: "success"}, nil
}
