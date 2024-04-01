package repository

import (
	"evaeats/user-service/internal/order/entity"

	"gorm.io/gorm"
)

// OrderItemRepositoryPostgres represents a PostgreSQL repository for managing OrderItem entities.
type OrderItemRepositoryPostgres struct {
	DB *gorm.DB
}

// NewOrderItemRepositoryPostgres creates a new instance of OrderItemRepositoryPostgres.
func NewOrderItemRepositoryPostgres(db *gorm.DB) *OrderItemRepositoryPostgres {
	return &OrderItemRepositoryPostgres{DB: db}
}

// Create adds a new OrderItem to the PostgreSQL database.
func (r *OrderItemRepositoryPostgres) Create(orderItem *entity.OrderItem) error {
	return r.DB.Create(orderItem).Error
}

// FindAll returns all OrderItems from the PostgreSQL database.
func (r *OrderItemRepositoryPostgres) FindAll() ([]*entity.OrderItem, error) {
	var orderItems []*entity.OrderItem
	if err := r.DB.Find(&orderItems).Error; err != nil {
		return nil, err
	}
	return orderItems, nil
}

// Update updates an existing OrderItem in the PostgreSQL database.
func (r *OrderItemRepositoryPostgres) Update(orderItem *entity.OrderItem) error {
	return r.DB.Save(orderItem).Error
}

// DeleteByID removes an OrderItem from the PostgreSQL database by ID.
func (r *OrderItemRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(&entity.OrderItem{}).Error
}

// GetByID retrieves an OrderItem from the PostgreSQL database by ID.
func (r *OrderItemRepositoryPostgres) GetByID(id string) (*entity.OrderItem, error) {
	var orderItem entity.OrderItem
	if err := r.DB.Where("id = ?", id).First(&orderItem).Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}
