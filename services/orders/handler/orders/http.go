package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Asad2730/Kitchen_Micro/common"
	"github.com/Asad2730/Kitchen_Micro/generated/orders"
	"github.com/Asad2730/Kitchen_Micro/services/orders/types"
)

type OrdersHttphandler struct {
	ordersService types.OrderService
}

func NewHttpHandler(ordersService types.OrderService) *OrdersHttphandler {
	return &OrdersHttphandler{ordersService: ordersService}
}

func (h *OrdersHttphandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttphandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	common.WriteJSON(w, http.StatusOK, res)
}
