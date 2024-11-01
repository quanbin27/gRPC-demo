package types

import (
	"context"
	"github.com/quanbin27/gRPC-demo/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
