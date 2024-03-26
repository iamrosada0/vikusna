package service

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type OrderItemService struct {
	OrderItemRepository repositories.OrderItemRepository
}

func NewOrderItemService(orderItemRepo repositories.OrderItemRepository) *OrderItemService {
	return &OrderItemService{
		OrderItemRepository: orderItemRepo,
	}
}

func (s *OrderItemService) CreateOrderItem(orderID, dishID string, quantity int) (*domain.OrderItem, error) {
	// Validate input data, if necessary

	// Create the order item in the database
	newOrderItem, err := s.OrderItemRepository.Insert(orderID, dishID, quantity)
	if err != nil {
		return nil, err
	}

	return newOrderItem, nil
}

func (s *OrderItemService) GetOrderItemByID(id string) (*domain.OrderItem, error) {
	// Fetch the order item by ID from the database
	orderItem, err := s.OrderItemRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}
