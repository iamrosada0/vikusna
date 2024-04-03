package repository

import (
	"evaeats/user-service/internal/payment/entity"

	"gorm.io/gorm"
)

type PaymentRepositoryPostgres struct {
	DB *gorm.DB
}

func NewPaymentRepositoryPostgres(db *gorm.DB) *PaymentRepositoryPostgres {
	return &PaymentRepositoryPostgres{DB: db}
}

func (r *PaymentRepositoryPostgres) Create(Payment *entity.Payment) error {
	return r.DB.Create(Payment).Error
}

func (r *PaymentRepositoryPostgres) FindAll() ([]*entity.Payment, error) {
	var Payments []*entity.Payment
	if err := r.DB.Find(&Payments).Error; err != nil {
		return nil, err
	}
	return Payments, nil
}

func (r *PaymentRepositoryPostgres) Update(Payment *entity.Payment) error {
	return r.DB.Save(Payment).Error
}

func (r *PaymentRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(entity.Payment{}).Error
}

func (r *PaymentRepositoryPostgres) GetByID(id string) (*entity.Payment, error) {
	var Payment entity.Payment
	if err := r.DB.Where("id = ?", id).First(&Payment).Error; err != nil {
		return nil, err
	}
	return &Payment, nil
}
