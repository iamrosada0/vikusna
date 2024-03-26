package services

import (
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
	// Validar os dados de entrada, se necessário

	// Criar o pedido no banco de dados
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
