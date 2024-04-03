package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// NotificationRepository define as operações disponíveis no repositório de notificações.
type NotificationRepository interface {
	Create(notification *Notification) error
	FindAll() ([]*Notification, error)
	Update(notification *Notification) error
	DeleteByID(id string) error
	GetByID(id string) (*Notification, error)
}

// Notification representa uma notificação enviada a um usuário.
type Notification struct {
	ID        string `json:"notification_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// NewNotification cria uma nova instância de Notification com os valores fornecidos.
func NewNotification(userID, message string) (*Notification, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Notification{
		ID:        id.String(),
		UserID:    userID,
		Message:   message,
		Timestamp: time.Now().Unix(),
	}, nil
}

// InMemoryNotificationRepository é uma implementação em memória da interface NotificationRepository.
type InMemoryNotificationRepository struct {
	Notifications map[string]*Notification
}

// NewInMemoryNotificationRepository cria um novo repositório de notificações em memória.
func NewInMemoryNotificationRepository() *InMemoryNotificationRepository {
	return &InMemoryNotificationRepository{
		Notifications: make(map[string]*Notification),
	}
}

// Create cria uma nova notificação.
func (r *InMemoryNotificationRepository) Create(notification *Notification) error {
	if _, exists := r.Notifications[notification.ID]; exists {
		return errors.New("notification already exists")
	}
	r.Notifications[notification.ID] = notification
	return nil
}

// DeleteByID exclui uma notificação pelo ID.
func (r *InMemoryNotificationRepository) DeleteByID(id string) error {
	if _, exists := r.Notifications[id]; !exists {
		return errors.New("notification not found")
	}
	delete(r.Notifications, id)
	return nil
}

// FindAll retorna todas as notificações.
func (r *InMemoryNotificationRepository) FindAll() ([]*Notification, error) {
	var allNotifications []*Notification
	for _, notification := range r.Notifications {
		allNotifications = append(allNotifications, notification)
	}
	return allNotifications, nil
}

// Update atualiza uma notificação.
func (r *InMemoryNotificationRepository) Update(notification *Notification) error {
	if _, exists := r.Notifications[notification.ID]; !exists {
		return errors.New("notification not found")
	}
	r.Notifications[notification.ID] = notification
	return nil
}

// GetByID retorna uma notificação pelo ID.
func (r *InMemoryNotificationRepository) GetByID(id string) (*Notification, error) {
	if notification, exists := r.Notifications[id]; exists {
		return notification, nil
	}
	return nil, errors.New("notification not found")
}

func (r *InMemoryNotificationRepository) GetAllByUserID(userID string) ([]*Notification, error) {
	var notifications []*Notification
	for _, notification := range r.Notifications {
		if notification.UserID == userID {
			notifications = append(notifications, notification)
		}
	}
	return notifications, nil
}
