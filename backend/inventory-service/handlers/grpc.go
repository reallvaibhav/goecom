package handlers

import (
	"context"
	"ecommerce/inventory-service/services"
	"ecommerce/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServiceServer
	service *services.ProductService
}

func NewInventoryServer(service *services.ProductService) *InventoryServer {
	return &InventoryServer{service: service}
}

func (s *InventoryServer) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	return s.service.CreateProduct(ctx, req)
}

func (s *InventoryServer) GetProductByID(ctx context.Context, req *proto.GetProductRequest) (*proto.ProductResponse, error) {
	return s.service.GetProductByID(ctx, req)
}

func (s *InventoryServer) UpdateProduct(ctx context.Context, req *proto.UpdateProductRequest) (*proto.ProductResponse, error) {
	return s.service.UpdateProduct(ctx, req)
}

func (s *InventoryServer) DeleteProduct(ctx context.Context, req *proto.DeleteProductRequest) (*proto.Empty, error) {
	return s.service.DeleteProduct(ctx, req)
}

func (s *InventoryServer) ListProducts(ctx context.Context, req *proto.ListProductsRequest) (*proto.ListProductsResponse, error) {
	return s.service.ListProducts(ctx, req)
}
