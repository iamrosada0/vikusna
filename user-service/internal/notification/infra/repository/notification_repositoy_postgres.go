package repository

import (
	"evaeats/user-service/internal/notification/entity"

	"gorm.io/gorm"
)

type NotificationRepositoryPostgres struct {
	DB *gorm.DB
}

func NewNotificationRepositoryPostgres(db *gorm.DB) *NotificationRepositoryPostgres {
	return &NotificationRepositoryPostgres{DB: db}
}

func (r *NotificationRepositoryPostgres) Create(notification *entity.Notification) error {
	return r.DB.Create(notification).Error
}

func (r *NotificationRepositoryPostgres) FindAll() ([]*entity.Notification, error) {
	var categories []*entity.Notification
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *NotificationRepositoryPostgres) Update(notification *entity.Notification) error {
	return r.DB.Save(notification).Error
}

func (r *NotificationRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(&entity.Notification{}).Error
}

func (r *NotificationRepositoryPostgres) GetByID(id string) (*entity.Notification, error) {
	var notification entity.Notification
	if err := r.DB.Where("id = ?", id).First(&notification).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}
