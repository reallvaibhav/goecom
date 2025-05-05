package handlers

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type Metrics struct {
	RequestCount map[string]int
	mu           sync.Mutex
}

func NewMetrics() *Metrics {
	return &Metrics{
		RequestCount: make(map[string]int),
	}
}

func (m *Metrics) Track() gin.HandlerFunc {
	return func(c *gin.Context) {
		m.mu.Lock()
		m.RequestCount[c.Request.URL.Path]++
		m.mu.Unlock()
		c.Next()
	}
}

func (m *Metrics) GetMetrics(c *gin.Context) {
	m.mu.Lock()
	defer m.mu.Unlock()
	c.JSON(200, m.RequestCount)
}
