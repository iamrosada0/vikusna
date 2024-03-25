package domain

import (
	"time"

	"github.com/google/uuid"
)

// Notificação representa uma notificação enviada a um usuário
type Notification struct {
	ID        string `json:"notification_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	UserID    string
	Message   string
	Timestamp int64
}

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
