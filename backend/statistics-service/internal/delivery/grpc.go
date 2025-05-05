package delivery

import (
	"context"
	"statistics-service/internal/usecase"
	"statistics-service/proto"
)

type Server struct {
	Usecase usecase.StatisticsUsecase
	proto.UnimplementedStatisticsServiceServer
}

func (s *Server) GetUserOrderStatistics(ctx context.Context, req *proto.UserOrderStatisticsRequest) (*proto.UserOrderStatisticsResponse, error) {
	// Fetch and return order statistics
	return &proto.UserOrderStatisticsResponse{}, nil
}

func (s *Server) GetUserStatistics(ctx context.Context, req *proto.UserStatisticsRequest) (*proto.UserStatisticsResponse, error) {
	// Fetch and return user statistics
	return &proto.UserStatisticsResponse{}, nil
}
