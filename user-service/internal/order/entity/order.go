package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(order *Order) error
	FindAll() ([]*Order, error)
	Update(order *Order) error
	DeleteByID(id string) error
	GetByID(id string) (*Order, error)
}

type Order struct {
	ID              string      `json:"order_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID      string      `json:"customer_id"`
	ChefID          string      `json:"chef_id"`
	Items           []OrderItem `json:"items"`
	Status          string      `json:"status"`
	OrderDate       string      `json:"order_date"`
	DriverID        string      `json:"driver_id" valid:"uuid" gorm:"type:uuid"`
	DeliveryAddress string      `json:"delivery_address"`
}

func NewOrder(customerID, chefID, status, driverID, deliveryAddress string, items []OrderItem) *Order {
	return &Order{
		ID:              uuid.New().String(),
		CustomerID:      customerID,
		ChefID:          chefID,
		Items:           items,
		Status:          status,
		OrderDate:       time.Now().Format(time.RFC3339),
		DriverID:        driverID,
		DeliveryAddress: deliveryAddress,
	}
}

func (d *Order) UpdateStatus(status string) {
	d.Status = status
}

type InMemoryOrderRepository struct {
	Orders map[string]*Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		Orders: make(map[string]*Order),
	}
}

func (r *InMemoryOrderRepository) Create(order *Order) error {
	if _, exists := r.Orders[order.ID]; exists {
		return errors.New("order already exists")
	}
	r.Orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepository) DeleteByID(id string) error {
	if _, exists := r.Orders[id]; !exists {
		return errors.New("order not found")
	}
	delete(r.Orders, id)
	return nil
}

func (r *InMemoryOrderRepository) FindAll() ([]*Order, error) {
	var allOrders []*Order
	for _, order := range r.Orders {
		allOrders = append(allOrders, order)
	}
	return allOrders, nil
}

func (r *InMemoryOrderRepository) Update(order *Order) error {
	if _, exists := r.Orders[order.ID]; !exists {
		return errors.New("order not found")
	}
	r.Orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepository) GetByID(id string) (*Order, error) {
	if order, exists := r.Orders[id]; exists {
		return order, nil
	}
	return nil, errors.New("order not found")
}
