package repository

import (
	"database/sql"
	"ecommerce/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *Repository) CreateOrder(req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		"INSERT INTO orders (user_id, status, total) VALUES (?, ?, ?)",
		req.UserId, "pending", 0.0,
	)
	if err != nil {
		return nil, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var total float64
	for _, item := range req.Items {
		// In production, fetch price from Inventory Service via gRPC
		price := 10.0 // Dummy price
		total += price * float64(item.Quantity)
		_, err := tx.Exec(
			"INSERT INTO order_items (order_id, product_id, quantity) VALUES (?, ?, ?)",
			orderID, item.ProductId, item.Quantity,
		)
		if err != nil {
			return nil, err
		}
	}

	_, err = tx.Exec("UPDATE orders SET total = ? WHERE order_id = ?", total, orderID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &proto.OrderResponse{
		OrderId: int32(orderID),
		UserId:  req.UserId,
		Items:   req.Items,
		Status:  "pending",
		Total:   total,
	}, nil
}

func (r *Repository) GetOrderByID(orderID int32) (*proto.OrderResponse, error) {
	row := r.db.QueryRow("SELECT order_id, user_id, status, total FROM orders WHERE order_id = ?", orderID)

	var o proto.OrderResponse
	err := row.Scan(&o.OrderId, &o.UserId, &o.Status, &o.Total)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query("SELECT product_id, quantity FROM order_items WHERE order_id = ?", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item proto.OrderItem
		if err := rows.Scan(&item.ProductId, &item.Quantity); err != nil {
			return nil, err
		}
		o.Items = append(o.Items, &item)
	}

	return &o, nil
}

func (r *Repository) UpdateOrderStatus(req *proto.UpdateOrderStatusRequest) (*proto.OrderResponse, error) {
	result, err := r.db.Exec("UPDATE orders SET status = ? WHERE order_id = ?", req.Status, req.OrderId)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return r.GetOrderByID(req.OrderId)
}

func (r *Repository) ListUserOrders(userID int32) ([]*proto.OrderResponse, error) {
	rows, err := r.db.Query("SELECT order_id, user_id, status, total FROM orders WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*proto.OrderResponse
	for rows.Next() {
		var o proto.OrderResponse
		if err := rows.Scan(&o.OrderId, &o.UserId, &o.Status, &o.Total); err != nil {
			return nil, err
		}
		order, err := r.GetOrderByID(o.OrderId)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
