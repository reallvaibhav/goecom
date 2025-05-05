package repository

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

type Repository interface {
    SaveOrderStatistics(userID string, totalOrders int, mostActiveTime string) error
    SaveUserStatistics(totalUsers, activeUsers int) error
}

type repository struct {
    db *sql.DB
}

func InitDB() (Repository, error) {
    connStr := "user=postgres password=yourpassword dbname=statistics sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    return &repository{db: db}, nil
}

func (r *repository) SaveOrderStatistics(userID string, totalOrders int, mostActiveTime string) error {
    // Save order statistics to the database
    return nil
}

func (r *repository) SaveUserStatistics(totalUsers, activeUsers int) error {
    // Save user statistics to the database
    return nil
}
```

// filepath: d:\Studies\Advance Programming\Assignment2_Kumar_Vaibhav\Ecom\backend\statistics-service\internal\usecase\statistics.go
package usecase

import "statistics-service/internal/repository"

type StatisticsUsecase interface {
    ProcessOrderCreated(data []byte)
    ProcessInventoryUpdated(data []byte)
}

type statisticsUsecase struct {
    repo repository.Repository
}

func NewStatisticsUsecase(repo repository.Repository) StatisticsUsecase {
    return &statisticsUsecase{repo: repo}
}

func (u *statisticsUsecase) ProcessOrderCreated(data []byte) {
    // Parse data and update statistics
}

func (u *statisticsUsecase) ProcessInventoryUpdated(data []byte) {
    // Parse data and update statistics
}
