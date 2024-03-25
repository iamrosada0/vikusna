package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	Insert(user_id, message string) (*domain.Notification, error)
	Find(id string) (*domain.Notification, error)
}

type NotificationRepositoryDb struct {
	Db *gorm.DB
}

func (repo NotificationRepositoryDb) Insert(user_id, message string) (*domain.Notification, error) {
	// Generate a new Notification with a generated ID
	newNotification, err := domain.NewNotification(user_id, message)
	if err != nil {
		return nil, err
	}

	// Insert the Notification into the database
	if err := repo.Db.Create(newNotification).Error; err != nil {
		return nil, err
	}

	return newNotification, nil
}

func (repo NotificationRepositoryDb) Find(id string) (*domain.Notification, error) {
	var notification domain.Notification
	if err := repo.Db.First(&notification, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("notification with ID %s not found", id)
		}
		return nil, err
	}
	return &notification, nil
}
