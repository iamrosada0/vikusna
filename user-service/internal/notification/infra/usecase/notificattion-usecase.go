package usecase

import (
	"evaeats/user-service/internal/notification/entity"
)

type DeleteNotificationInputDto struct {
	ID string `json:"notification_id" valid:"uuid"`
}
type UpdateNotificationInputDto struct {
	ID      string `json:"notification_id" valid:"uuid"`
	Message string `json:"message"`
}
type NotificationOutputDto struct {
	ID        string `json:"notification_id" valid:"uuid"`
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type CreateNotificationInputDto struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type CreateNotificationOutputDto struct {
	ID        string `json:"notification_id" valid:"uuid"`
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type CreateNotificationUseCase struct {
	NotificationRepository entity.NotificationRepository
}

func NewCreateNotificationUseCase(notificationRepository entity.NotificationRepository) *CreateNotificationUseCase {
	return &CreateNotificationUseCase{NotificationRepository: notificationRepository}
}

// Método para criar uma nova notificação.
func (u *CreateNotificationUseCase) Execute(input CreateNotificationInputDto) (*CreateNotificationOutputDto, error) {
	newNotification, err := entity.NewNotification(input.UserID, input.Message)
	if err != nil {
		return nil, err
	}

	err = u.NotificationRepository.Create(newNotification)
	if err != nil {
		return nil, err
	}

	output := &CreateNotificationOutputDto{
		ID:        newNotification.ID,
		UserID:    newNotification.UserID,
		Message:   newNotification.Message,
		Timestamp: newNotification.Timestamp,
	}

	return output, nil
}

// Caso de uso para excluir uma notificação.
type DeleteNotificationUseCase struct {
	NotificationRepository entity.NotificationRepository
}

// Função para criar uma nova instância de DeleteNotificationUseCase.
func NewDeleteNotificationUseCase(notificationRepository entity.NotificationRepository) *DeleteNotificationUseCase {
	return &DeleteNotificationUseCase{NotificationRepository: notificationRepository}
}

// Método para excluir uma notificação.
func (u *DeleteNotificationUseCase) Execute(input DeleteNotificationInputDto) error {
	err := u.NotificationRepository.DeleteByID(input.ID)
	if err != nil {
		return err
	}
	return nil
}

// Caso de uso para atualizar uma notificação.
type UpdateNotificationUseCase struct {
	NotificationRepository entity.NotificationRepository
}

// Função para criar uma nova instância de UpdateNotificationUseCase.
func NewUpdateNotificationUseCase(notificationRepository entity.NotificationRepository) *UpdateNotificationUseCase {
	return &UpdateNotificationUseCase{NotificationRepository: notificationRepository}
}

// Método para atualizar uma notificação.
func (u *UpdateNotificationUseCase) Execute(input UpdateNotificationInputDto) error {
	notification, err := u.NotificationRepository.GetByID(input.ID)
	if err != nil {
		return err
	}
	notification.Message = input.Message
	err = u.NotificationRepository.Update(notification)
	if err != nil {
		return err
	}
	return nil
}

// Caso de uso para buscar notificações.
type GetNotificationsUseCase struct {
	NotificationRepository entity.NotificationRepository
}

// Função para criar uma nova instância de GetNotificationsUseCase.
func NewGetNotificationsUseCase(notificationRepository entity.NotificationRepository) *GetNotificationsUseCase {
	return &GetNotificationsUseCase{NotificationRepository: notificationRepository}
}

// Método para buscar notificações.
func (u *GetNotificationsUseCase) Execute() ([]*NotificationOutputDto, error) {
	notifications, err := u.NotificationRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*NotificationOutputDto
	for _, notification := range notifications {
		output = append(output, &NotificationOutputDto{
			ID:        notification.ID,
			UserID:    notification.UserID,
			Message:   notification.Message,
			Timestamp: notification.Timestamp,
		})
	}

	return output, nil
}
