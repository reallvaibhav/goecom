package services

import (
	"context"
	"ecommerce/inventory-service/repository"
	"ecommerce/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	repo *repository.Repository
}

func NewProductService(repo *repository.Repository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	if req.Name == "" || req.Category == "" || req.Stock < 0 || req.Price < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product details")
	}
	return s.repo.CreateProduct(req)
}

func (s *ProductService) GetProductByID(ctx context.Context, req *proto.GetProductRequest) (*proto.ProductResponse, error) {
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product ID")
	}
	return s.repo.GetProductByID(req.Id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.ProductResponse, error) {
	if req.Id <= 0 || req.Name == "" || req.Category == "" || req.Stock < 0 || req.Price < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product details")
	}
	return s.repo.UpdateProduct(req)
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.Empty, error) {
	if req.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product ID")
	}
	err := s.repo.DeleteProduct(req.Id)
	return &proto.Empty{}, err
}

func (s *ProductService) ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*proto.ListProductsResponse, error) {
	products, err := s.repo.ListProducts()
	if err != nil {
		return nil, err
	}
	return &proto.ListProductsResponse{Products: products}, nil
}
