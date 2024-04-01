package usecase

import (
	"evaeats/user-service/internal/notification/entity"
)

// Aqui está a definição de uma estrutura de entrada para a operação de exclusão de notificação.
type DeleteNotificationInputDto struct {
	ID string `json:"notification_id" valid:"uuid"`
}

// Aqui está a definição de uma estrutura de entrada para a operação de atualização de notificação.
type UpdateNotificationInputDto struct {
	ID      string `json:"notification_id" valid:"uuid"`
	Message string `json:"message"`
}

// Aqui está a definição de uma estrutura de saída para as operações de busca, exclusão e atualização de notificações.
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
	ID        string `json:"notification_id" valid:"uuid" gorm:"type:uuid;primary_key"`
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

// Aqui está a definição do caso de uso para a operação de exclusão de notificação.
type DeleteNotificationUseCase struct {
	NotificationRepository entity.NotificationRepository
}

// Aqui está a definição do caso de uso para a operação de atualização de notificação.
type UpdateNotificationUseCase struct {
	NotificationRepository entity.NotificationRepository
}

// Aqui está a implementação do método Execute para o caso de uso de exclusão de notificação.
func (u *DeleteNotificationUseCase) Execute(input DeleteNotificationInputDto) error {
	err := u.NotificationRepository.DeleteByID(input.ID)
	if err != nil {
		return err
	}
	return nil
}

// Aqui está a implementação do método Execute para o caso de uso de atualização de notificação.
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

// Aqui está a definição do caso de uso para a operação de busca de notificações.
type GetNotificationsUseCase struct {
	NotificationRepository entity.NotificationRepository
}

// Aqui está a implementação do método Execute para o caso de uso de busca de notificações.
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
