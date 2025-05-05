package services

import (
	"context"
	"ecommerce/order-service/repository"
	"ecommerce/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
	repo *repository.Repository
	proto.UnimplementedOrderServiceServer
}

func NewOrderService(repo *repository.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	if req.UserId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid user ID")
	}
	if len(req.Items) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty items")
	}
	for _, item := range req.Items {
		if item.ProductId <= 0 || item.Quantity <= 0 {
			return nil, status.Error(codes.InvalidArgument, "invalid product ID or quantity")
		}
	}
	return s.repo.CreateOrder(req)
}

func (s *OrderService) GetOrderByID(ctx context.Context, req *proto.GetOrderRequest) (*proto.OrderResponse, error) {
	if req.OrderId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid order ID")
	}
	return s.repo.GetOrderByID(req.OrderId)
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, req *proto.UpdateOrderStatusRequest) (*proto.OrderResponse, error) {
	if req.OrderId <= 0 || req.Status == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid order ID or status")
	}
	return s.repo.UpdateOrderStatus(req)
}

func (s *OrderService) ListUserOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	if req.UserId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid user ID")
	}
	orders, err := s.repo.ListUserOrders(req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto.ListOrdersResponse{Orders: orders}, nil
}
