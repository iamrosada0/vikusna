package entity

import (
	"errors"
	"math/rand"
)

type PaymentRepository interface {
	Create(payment *Payment) error
	FindAll() ([]*Payment, error)
	Update(payment *Payment) error
	DeleteByID(id uint) error
	GetByID(id uint) (*Payment, error)
}

type Payment struct {
	ID    uint
	Name  string
	Email string
}

func NewPayment(name, email string) *Payment {
	return &Payment{
		ID:    uint(rand.Uint32()),
		Name:  name,
		Email: email,
	}
}

func (d *Payment) Update(name, email string) {
	d.Name = name
	d.Email = email
}

type InMemoryPaymentRepository struct {
	Payments map[string]*Payment
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{
		Payments: make(map[string]*Payment),
	}
}

func (r *InMemoryPaymentRepository) DeleteByID(id string) error {
	if _, exists := r.Payments[id]; !exists {
		return errors.New("Payment not found")
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
