package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type PaymentTransactionRepository interface {
	Insert(order_id, payment_method string, amount float64) (*domain.PaymentTransaction, error)
	Find(id string) (*domain.PaymentTransaction, error)
}

type PaymentTransactionRepositoryDb struct {
	Db *gorm.DB
}

func (repo PaymentTransactionRepositoryDb) Insert(order_id, payment_method string, amount float64) (*domain.PaymentTransaction, error) {
	// Generate a new PaymentTransaction with a generated ID
	newPaymentTransaction, err := domain.NewPaymentTransaction(order_id, payment_method, amount)
	if err != nil {
		return nil, err
	}

	// Insert the PaymentTransaction into the database
	if err := repo.Db.Create(newPaymentTransaction).Error; err != nil {
		return nil, err
	}

	return newPaymentTransaction, nil
}

func (repo PaymentTransactionRepositoryDb) Find(id string) (*domain.PaymentTransaction, error) {
	var paymentTransaction domain.PaymentTransaction
	if err := repo.Db.First(&paymentTransaction, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("payment with ID %s not found", id)
		}
		return nil, err
	}
	return &paymentTransaction, nil
}
