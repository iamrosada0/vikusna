package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	Insert(orderID, paymentMethod, paymentStatus, paymentDate, userID string, amount float64) (*domain.Invoice, error)
	Find(id string) (*domain.Invoice, error)
}

type InvoiceRepositoryDb struct {
	Db *gorm.DB
}

func NewInvoiceRepository(Db *gorm.DB) *InvoiceRepositoryDb {
	return &InvoiceRepositoryDb{Db: Db}
}

func (repo *InvoiceRepositoryDb) Insert(orderID, paymentMethod, paymentStatus, paymentDate, userID string, amount float64) (*domain.Invoice, error) {
	// Generate a new Invoice with a generated ID
	newInvoice := &domain.Invoice{
		OrderID:        orderID,
		Payment_method: paymentMethod,
		Payment_status: paymentStatus,
		Payment_date:   paymentDate,
		UserID:         userID,
		Amount:         amount,
	}

	// Insert the Invoice into the database
	if err := repo.Db.Create(newInvoice).Error; err != nil {
		return nil, err
	}

	return newInvoice, nil
}

func (repo *InvoiceRepositoryDb) Find(id string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	if err := repo.Db.First(&invoice, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Invoice with ID %s not found", id)
		}
		return nil, err
	}
	return &invoice, nil
}
