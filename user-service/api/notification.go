package api

import (
	"evaeats/user-service/internal/notification/infra/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandlers struct {
	CreateNotificationUseCase *usecase.CreateNotificationUseCase
	GetNotificationsUseCase   *usecase.GetNotificationsUseCase
	DeleteNotificationUseCase *usecase.DeleteNotificationUseCase
	UpdateNotificationUseCase *usecase.UpdateNotificationUseCase
}

func NewNotificationHandlers(
	createNotificationUseCase *usecase.CreateNotificationUseCase,
	getNotificationsUseCase *usecase.GetNotificationsUseCase,
	deleteNotificationUseCase *usecase.DeleteNotificationUseCase,
	updateNotificationUseCase *usecase.UpdateNotificationUseCase,
) *NotificationHandlers {
	return &NotificationHandlers{
		CreateNotificationUseCase: createNotificationUseCase,
		GetNotificationsUseCase:   getNotificationsUseCase,
		DeleteNotificationUseCase: deleteNotificationUseCase,
		UpdateNotificationUseCase: updateNotificationUseCase,
	}
}

func (nh *NotificationHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		notifications := api.Group("/notifications")
		{
			notifications.POST("/", nh.CreateNotificationHandler)
			notifications.GET("/", nh.GetNotificationsHandler)
			notifications.DELETE("/:id", nh.DeleteNotificationHandler)
			notifications.PUT("/:id", nh.UpdateNotificationHandler)
		}
	}
}

func (nh *NotificationHandlers) CreateNotificationHandler(c *gin.Context) {
	var input usecase.CreateNotificationInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := nh.CreateNotificationUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (nh *NotificationHandlers) GetNotificationsHandler(c *gin.Context) {
	output, err := nh.GetNotificationsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (nh *NotificationHandlers) DeleteNotificationHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.DeleteNotificationInputDto{ID: id}
	err := nh.DeleteNotificationUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}

func (nh *NotificationHandlers) UpdateNotificationHandler(c *gin.Context) {
	var input usecase.UpdateNotificationInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := nh.UpdateNotificationUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notification updated successfully"})
}
