package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(dbPath string) (*Repository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			order_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			status TEXT NOT NULL,
			total REAL NOT NULL
		);
		CREATE TABLE IF NOT EXISTS order_items (
			order_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL,
			FOREIGN KEY (order_id) REFERENCES orders(order_id)
		)
	`)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Close() {
	if err := r.db.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
