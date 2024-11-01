package service

import (
	"context"
	"github.com/quanbin27/gRPC-demo/services/common/genproto/orders"
)

var orderlists = make([]*orders.Order, 0)

type OrderService struct {
	//store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	orderlists = append(orderlists, order)
	return nil
}
func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return orderlists
}
