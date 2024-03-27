package services

import (
	"errors"
	"evaeats/application/repositories"
	"evaeats/domain"
)

type OrderService struct {
	OrderRepository repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) *OrderService {
	return &OrderService{
		OrderRepository: orderRepo,
	}
}

func (s *OrderService) CreateOrder(customerID, chefID, status, address string, items []domain.OrderItem) (*domain.Order, error) {
	// Validate input data
	if customerID == "" {
		return nil, errors.New("customerID cannot be empty")
	}
	if chefID == "" {
		return nil, errors.New("chefID cannot be empty")
	}
	if status != "PENDING" && status != "ACCEPTED" && status != "REJECTED" {
		return nil, errors.New("invalid status")
	}
	if address == "" {
		return nil, errors.New("address cannot be empty")
	}
	if len(items) == 0 {
		return nil, errors.New("order must contain at least one item")
	}
	// Additional validation logic can be added as needed

	// Create the order in the database
	newOrder, err := s.OrderRepository.Insert(customerID, chefID, status, address, items)
	if err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (s *OrderService) GetOrderByID(id string) (*domain.Order, error) {
	// Buscar o pedido pelo ID no banco de dados
	order, err := s.OrderRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
