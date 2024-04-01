package api

import (
	"evaeats/user-service/internal/payment/infra/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandlers struct {
	ProcessPaymentUseCase *usecase.CreatePaymentUseCase
}

func NewPaymentHandlers(processPaymentUseCase *usecase.CreatePaymentUseCase) *PaymentHandlers {
	return &PaymentHandlers{
		ProcessPaymentUseCase: processPaymentUseCase,
	}
}

func (ph *PaymentHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		payment := api.Group("/payment")
		{
			payment.POST("/", ph.ProcessPaymentHandler)
		}
	}
}

func (ph *PaymentHandlers) ProcessPaymentHandler(c *gin.Context) {
	var input usecase.CreatePaymentInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := ph.ProcessPaymentUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
