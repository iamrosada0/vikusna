package api

import (
	"evaeats/user-service/internal/order/infra/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderItemHandlers struct {
	CreateOrderItemUseCase *usecase.CreateOrderItemUseCase
	// ListOrderItemesUseCase  *usecase.GetAllOrdersUseCase
	DeleteOrderItemUseCase  *usecase.DeleteOrderItemUseCase
	GetOrderItemByIDUseCase *usecase.GetOrderItemUseCase
	UpdateOrderItemUseCase  *usecase.UpdateOrderItemUseCase
}

func NewOrderItemHandlers(
	createOrderItemUseCase *usecase.CreateOrderItemUseCase,
	deleteOrderItemUseCase *usecase.DeleteOrderItemUseCase,
	getOrderItemByIDUseCase *usecase.GetOrderItemUseCase,
	updateOrderItemUseCase *usecase.UpdateOrderItemUseCase,
) *OrderItemHandlers {
	return &OrderItemHandlers{
		CreateOrderItemUseCase:  createOrderItemUseCase,
		DeleteOrderItemUseCase:  deleteOrderItemUseCase,
		GetOrderItemByIDUseCase: getOrderItemByIDUseCase,
		UpdateOrderItemUseCase:  updateOrderItemUseCase,
	}
}

func (dh *OrderItemHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		OrderItemes := api.Group("/order-items")
		{
			OrderItemes.POST("/", dh.CreateOrderItemHandler)
			OrderItemes.DELETE("/:id", dh.DeleteOrderItemHandler)
			OrderItemes.GET("/:id", dh.GetOrderItemByIDHandler)
			OrderItemes.PUT("/:id", dh.UpdateOrderItemHandler)
		}
	}
}

func (dh *OrderItemHandlers) CreateOrderItemHandler(c *gin.Context) {
	var input usecase.CreateOrderItemInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := dh.CreateOrderItemUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

// func (dh *OrderItemHandlers) ListOrderItemesHandler(c *gin.Context) {
// 	output, err := dh.ListOrderItemesUseCase.Execute()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, output)
// }

func (dh *OrderItemHandlers) DeleteOrderItemHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.DeleteOrderItemInputDto{ID: id}
	output := dh.DeleteOrderItemUseCase.Execute(input)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, output)
}

func (dh *OrderItemHandlers) GetOrderItemByIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.GetOrderByIDInputDto{ID: id}
	output, err := dh.GetOrderItemByIDUseCase.Execute(usecase.GetOrderItemInputDto(input))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dh *OrderItemHandlers) UpdateOrderItemHandler(c *gin.Context) {
	id := c.Param("id")
	var input usecase.UpdateOrderItemInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id

	output, err := dh.UpdateOrderItemUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
