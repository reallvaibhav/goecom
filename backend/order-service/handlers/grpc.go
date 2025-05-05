package handlers

import (
	"context"
	"ecommerce/order-service/services"
	"ecommerce/proto"
)

type OrderServer struct {
	proto.UnimplementedOrderServiceServer
	service *services.OrderService
}

func NewOrderServer(service *services.OrderService) *OrderServer {
	return &OrderServer{service: service}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	return s.service.CreateOrder(ctx, req)
}

func (s *OrderServer) GetOrderByID(ctx context.Context, req *proto.GetOrderRequest) (*proto.OrderResponse, error) {
	return s.service.GetOrderByID(ctx, req)
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *proto.UpdateOrderStatusRequest) (*proto.OrderResponse, error) {
	return s.service.UpdateOrderStatus(ctx, req)
}

func (s *OrderServer) ListUserOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	return s.service.ListUserOrders(ctx, req)
}
