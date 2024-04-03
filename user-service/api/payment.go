package api

import (
	"evaeats/user-service/internal/payment/infra/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandlers struct {
	CreatePaymentUseCase  *usecase.CreatePaymentUseCase
	ListPaymentsUseCase   *usecase.GetAllPaymentsUseCase
	GetPaymentByIDUseCase *usecase.GetPaymentByIDUseCase
	UpdatePaymentUseCase  *usecase.UpdatePaymentUseCase
	DeletePaymentUseCase  *usecase.DeletePaymentUseCase
}

func NewPaymentHandlers(
	createPaymentUseCase *usecase.CreatePaymentUseCase,
	listPaymentsUseCase *usecase.GetAllPaymentsUseCase,
	getPaymentByIDUseCase *usecase.GetPaymentByIDUseCase,
	updatePaymentUseCase *usecase.UpdatePaymentUseCase,
	deletePaymentUseCase *usecase.DeletePaymentUseCase,
) *PaymentHandlers {
	return &PaymentHandlers{
		CreatePaymentUseCase:  createPaymentUseCase,
		ListPaymentsUseCase:   listPaymentsUseCase,
		GetPaymentByIDUseCase: getPaymentByIDUseCase,
		UpdatePaymentUseCase:  updatePaymentUseCase,
		DeletePaymentUseCase:  deletePaymentUseCase,
	}
}

func (ph *PaymentHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		payment := api.Group("/payment")
		{
			payment.POST("/", ph.CreatePaymentHandler)
			payment.GET("/", ph.ListPaymentsHandler)
			payment.GET("/:id", ph.GetPaymentByIDHandler)
			payment.PUT("/:id", ph.UpdatePaymentHandler)
			payment.DELETE("/:id", ph.DeletePaymentHandler)
		}
	}
}

func (ph *PaymentHandlers) CreatePaymentHandler(c *gin.Context) {
	var input usecase.CreatePaymentInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := ph.CreatePaymentUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (ph *PaymentHandlers) ListPaymentsHandler(c *gin.Context) {
	output, err := ph.ListPaymentsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (ph *PaymentHandlers) GetPaymentByIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.GetPaymentByIDInputDto{ID: id}
	output, err := ph.GetPaymentByIDUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (ph *PaymentHandlers) UpdatePaymentHandler(c *gin.Context) {
	id := c.Param("id")
	var input usecase.UpdatePaymentInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id

	output, err := ph.UpdatePaymentUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (ph *PaymentHandlers) DeletePaymentHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.DeletePaymentInputDto{ID: id}
	output, err := ph.DeletePaymentUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
