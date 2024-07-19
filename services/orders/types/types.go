package types

import (
	"context"

	"github.com/Asad2730/Kitchen_Micro/generated/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
