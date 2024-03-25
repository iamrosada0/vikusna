package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Insert(order_id, dish_id string, quantity int) (*domain.OrderItem, error)
	Find(id string) (*domain.OrderItem, error)
}

type OrderItemRepositoryDb struct {
	Db *gorm.DB
}

func (repo OrderItemRepositoryDb) Insert(order_id, dish_id string, quantity int) (*domain.OrderItem, error) {
	// Generate a new OrderItem with a generated ID
	newOrderItem, err := domain.NewOrderItem(order_id, dish_id, quantity)
	if err != nil {
		return nil, err
	}

	// Insert the OrderItem into the database
	if err := repo.Db.Create(newOrderItem).Error; err != nil {
		return nil, err
	}

	return newOrderItem, nil
}

func (repo OrderItemRepositoryDb) Find(id string) (*domain.OrderItem, error) {
	var OrderItem domain.OrderItem
	if err := repo.Db.First(&OrderItem, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("dish category with ID %s not found", id)
		}
		return nil, err
	}
	return &OrderItem, nil
}
