package entity

import (
	"errors"
	"math/rand"
)

type OrderRepository interface {
	Create(order *Order) error
	FindAll() ([]*Order, error)
	Update(order *Order) error
	DeleteByID(id uint) error
	GetByID(id uint) (*Order, error)
}

type Order struct {
	ID    uint
	Name  string
	Email string
}

func NewOrder(name, email string) *Order {
	return &Order{
		ID:    uint(rand.Uint32()),
		Name:  name,
		Email: email,
	}
}

func (d *Order) Update(name, email string) {
	d.Name = name
	d.Email = email
}

type InMemoryOrderRepository struct {
	Orders map[string]*Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		Orders: make(map[string]*Order),
	}
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
