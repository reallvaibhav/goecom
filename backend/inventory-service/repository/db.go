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
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			category TEXT NOT NULL,
			stock INTEGER NOT NULL,
			price REAL NOT NULL
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
