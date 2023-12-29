package services

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"sync"
	"wb_0_service/app/models"
)

type OrderStorageService struct {
	mutex sync.RWMutex
	cache map[string]models.Order
	db    *sqlx.DB
}

func NewOrderStorageService(db *sqlx.DB) *OrderStorageService {
	return &OrderStorageService{
		mutex: sync.RWMutex{},
		cache: make(map[string]models.Order),
		db:    db,
	}
}

func (s *OrderStorageService) AddOrder(order models.Order) error {
	bytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("INSERT INTO orders (order_id, json) VALUES ($1, $2)", order.OrderUid, bytes); err != nil {
		return err
	}

	s.mutex.Lock()
	s.cache[order.OrderUid] = order
	s.mutex.Unlock()

	return nil
}

func (s *OrderStorageService) GetOrderById(id string) (*models.Order, error) {
	if order, ok := s.cache[id]; ok {
		return &order, nil
	}

	s.mutex.RLock()
	var jason []byte
	s.mutex.RUnlock()

	err := s.db.Get(&jason, "SELECT json FROM orders WHERE order_id=$1", id)
	if err != nil {
		return nil, err
	}

	order := models.Order{}
	_ = json.Unmarshal(jason, &order)

	return &order, nil
}

func (s *OrderStorageService) RestoreCache() error {
	var orders []struct {
		OrderID string `db:"order_id"`
		Json    []byte `db:"json"`
	}

	err := s.db.Select(&orders, "SELECT order_id, json FROM orders")
	if err != nil {
		return err
	}

	s.mutex.Lock()
	for _, orderData := range orders {
		order := models.Order{}
		_ = json.Unmarshal(orderData.Json, &order)
		s.cache[orderData.OrderID] = order
	}
	s.mutex.Unlock()

	return nil
}
