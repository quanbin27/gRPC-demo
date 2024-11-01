package handler

import (
	"context"
	"github.com/quanbin27/gRPC-demo/services/common/genproto/orders"
	"github.com/quanbin27/gRPC-demo/services/orders/types"

	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	//service injection
	ordersService types.OrderService
	//unimqlemented service server
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	// register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

//func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
//	o := h.ordersService.GetOrders(ctx)
//	res := &orders.GetOrderResponse{
//		Orders: o,
//	}
//
//	return res, nil
//}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    1,
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}
	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
