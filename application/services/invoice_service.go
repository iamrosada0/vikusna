package services

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type InvoiceService struct {
	InvoiceRepository repositories.InvoiceRepository
}

func NewInvoiceService(invoiceRepo repositories.InvoiceRepository) *InvoiceService {
	return &InvoiceService{
		InvoiceRepository: invoiceRepo,
	}
}

func (s *InvoiceService) CreateInvoice(orderID, paymentMethod, paymentStatus, paymentDate, userID string, amount float64) (*domain.Invoice, error) {
	// Create a new invoice in the repository
	newInvoice, err := s.InvoiceRepository.Insert(orderID, paymentMethod, paymentStatus, paymentDate, userID, amount)
	if err != nil {
		return nil, err
	}

	return newInvoice, nil
}

func (s *InvoiceService) GetInvoiceByID(id string) (*domain.Invoice, error) {
	// Retrieve the invoice by ID from the repository
	invoice, err := s.InvoiceRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}
