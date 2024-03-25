package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Insert(customer_id, chef_id, status, address string, items []domain.OrderItem) (*domain.Order, error)
	Find(id string) (*domain.Order, error)
}

type OrderRepositoryDb struct {
	Db *gorm.DB
}

func (repo OrderRepositoryDb) Insert(customer_id, chef_id, status, address string, items []domain.OrderItem) (*domain.Order, error) {
	// Generate a new Order with a generated ID
	newOrder, err := domain.NewOrder(customer_id, chef_id, items, status, address)
	if err != nil {
		return nil, err
	}

	// Insert the Order into the database
	if err := repo.Db.Create(newOrder).Error; err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (repo OrderRepositoryDb) Find(id string) (*domain.Order, error) {
	var order domain.Order
	if err := repo.Db.First(&order, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Order with ID %s not found", id)
		}
		return nil, err
	}
	return &order, nil
}
