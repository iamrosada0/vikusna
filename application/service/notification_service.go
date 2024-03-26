package service

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type NotificationService struct {
	NotificationRepository repositories.NotificationRepository
}

func NewNotificationService(notificationRepo repositories.NotificationRepository) *NotificationService {
	return &NotificationService{
		NotificationRepository: notificationRepo,
	}
}

func (s *NotificationService) CreateNotification(userID, message string) (*domain.Notification, error) {
	// Validate input data, if necessary

	// Create the notification in the database
	newNotification, err := s.NotificationRepository.Insert(userID, message)
	if err != nil {
		return nil, err
	}

	return newNotification, nil
}

func (s *NotificationService) GetNotificationByID(id string) (*domain.Notification, error) {
	// Fetch the notification by ID from the database
	notification, err := s.NotificationRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return notification, nil
}
