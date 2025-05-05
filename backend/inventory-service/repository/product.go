package repository

import (
	"database/sql"
	"ecommerce/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *Repository) CreateProduct(req *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	result, err := r.db.Exec(
		"INSERT INTO products (name, category, stock, price) VALUES (?, ?, ?, ?)",
		req.Name, req.Category, req.Stock, req.Price,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &proto.ProductResponse{
		Id:       int32(id),
		Name:     req.Name,
		Category: req.Category,
		Stock:    req.Stock,
		Price:    req.Price,
	}, nil
}

func (r *Repository) GetProductByID(id int32) (*proto.ProductResponse, error) {
	row := r.db.QueryRow("SELECT id, name, category, stock, price FROM products WHERE id = ?", id)

	var p proto.ProductResponse
	err := row.Scan(&p.Id, &p.Name, &p.Category, &p.Stock, &p.Price)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *Repository) UpdateProduct(req *proto.UpdateProductRequest) (*proto.ProductResponse, error) {
	result, err := r.db.Exec(
		"UPDATE products SET name = ?, category = ?, stock = ?, price = ? WHERE id = ?",
		req.Name, req.Category, req.Stock, req.Price, req.Id,
	)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound, "product not found")
	}

	return &proto.ProductResponse{
		Id:       req.Id,
		Name:     req.Name,
		Category: req.Category,
		Stock:    req.Stock,
		Price:    req.Price,
	}, nil
}

func (r *Repository) DeleteProduct(id int32) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return status.Error(codes.NotFound, "product not found")
	}

	return nil
}

func (r *Repository) ListProducts() ([]*proto.ProductResponse, error) {
	rows, err := r.db.Query("SELECT id, name, category, stock, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*proto.ProductResponse
	for rows.Next() {
		var p proto.ProductResponse
		if err := rows.Scan(&p.Id, &p.Name, &p.Category, &p.Stock, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}
