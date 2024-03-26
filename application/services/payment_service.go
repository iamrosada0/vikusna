package services

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type PaymentTransactionService struct {
	PaymentTransactionRepository repositories.PaymentTransactionRepository
}

func NewPaymentTransactionService(paymentRepo repositories.PaymentTransactionRepository) *PaymentTransactionService {
	return &PaymentTransactionService{
		PaymentTransactionRepository: paymentRepo,
	}
}

func (s *PaymentTransactionService) CreatePaymentTransaction(orderID, paymentMethod string, amount float64) (*domain.PaymentTransaction, error) {
	// Validar os dados de entrada, se necessário

	// Criar a transação de pagamento no banco de dados
	newPaymentTransaction, err := s.PaymentTransactionRepository.Insert(orderID, paymentMethod, amount)
	if err != nil {
		return nil, err
	}

	return newPaymentTransaction, nil
}

func (s *PaymentTransactionService) GetPaymentTransactionByID(id string) (*domain.PaymentTransaction, error) {
	// Buscar a transação de pagamento pelo ID no banco de dados
	paymentTransaction, err := s.PaymentTransactionRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return paymentTransaction, nil
}
