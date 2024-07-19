package handler

import (
	"context"

	"github.com/Asad2730/Kitchen_Micro/generated/orders"
	"github.com/Asad2730/Kitchen_Micro/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcHandler(gRpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{orderService: orderService}

	orders.RegisterOrderServiceServer(gRpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{Status: "success"}, nil
}
