package repository

import (
	"evaeats/user-service/internal/order/entity"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrderRepositoryPostgres struct {
	DB *gorm.DB
}

func NewOrderRepositoryPostgres(db *gorm.DB) *OrderRepositoryPostgres {
	return &OrderRepositoryPostgres{DB: db}
}

func (r *OrderRepositoryPostgres) Create(order *entity.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepositoryPostgres) FindAll() ([]*entity.Order, error) {
	var Orders []*entity.Order
	if err := r.DB.Find(&Orders).Error; err != nil {
		return nil, err
	}
	return Orders, nil
}

func (r *OrderRepositoryPostgres) Update(order *entity.Order) error {
	return r.DB.Save(order).Error
}

func (r *OrderRepositoryPostgres) DeleteByID(id uint) error {
	return r.DB.Where("id = ?", id).Delete(entity.Order{}).Error
}

func (r *OrderRepositoryPostgres) GetByID(id uint) (*entity.Order, error) {
	var order entity.Order
	if err := r.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
