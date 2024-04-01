package entity

import (
	"errors"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(payment *Payment) error
	FindAll() ([]*Payment, error)
	Update(payment *Payment) error
	DeleteByID(id string) error
	GetByID(id string) (*Payment, error)
}

type Payment struct {
	ID            string  `json:"payment_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

func NewPaymentTransaction(orderID, paymentMethod string, amount float64) (*Payment, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Payment{
		ID:            id.String(),
		OrderID:       orderID,
		Amount:        amount,
		PaymentMethod: paymentMethod,
	}, nil
}

type InMemoryPaymentRepository struct {
	Payments map[string]*Payment
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{
		Payments: make(map[string]*Payment),
	}
}

func (r *InMemoryPaymentRepository) Create(payment *Payment) error {
	if _, exists := r.Payments[payment.ID]; exists {
		return errors.New("payment already exists")
	}
	r.Payments[payment.ID] = payment
	return nil
}

func (r *InMemoryPaymentRepository) DeleteByID(id string) error {
	if _, exists := r.Payments[id]; !exists {
		return errors.New("payment not found")
	}
	delete(r.Payments, id)
	return nil
}

func (r *InMemoryPaymentRepository) FindAll() ([]*Payment, error) {
	var allPayments []*Payment
	for _, payment := range r.Payments {
		allPayments = append(allPayments, payment)
	}
	return allPayments, nil
}

func (r *InMemoryPaymentRepository) Update(payment *Payment) error {
	if _, exists := r.Payments[payment.ID]; !exists {
		return errors.New("payment not found")
	}
	r.Payments[payment.ID] = payment
	return nil
}

func (r *InMemoryPaymentRepository) GetByID(id string) (*Payment, error) {
	if payment, exists := r.Payments[id]; exists {
		return payment, nil
	}
	return nil, errors.New("payment not found")
}
