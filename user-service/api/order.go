package api

import (
	"evaeats/user-service/internal/order/infra/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandlers struct {
	CreateOrderUseCase  *usecase.CreateOrderUseCase
	ListOrderesUseCase  *usecase.GetAllOrdersUseCase
	DeleteOrderUseCase  *usecase.DeleteOrderUseCase
	GetOrderByIDUseCase *usecase.GetOrderByIDUseCase
	UpdateOrderUseCase  *usecase.UpdateOrderUseCase
}

func NewOrderHandlers(
	createOrderUseCase *usecase.CreateOrderUseCase,
	listOrderesUseCase *usecase.GetAllOrdersUseCase,
	deleteOrderUseCase *usecase.DeleteOrderUseCase,
	getOrderByIDUseCase *usecase.GetOrderByIDUseCase,
	updateOrderUseCase *usecase.UpdateOrderUseCase,
) *OrderHandlers {
	return &OrderHandlers{
		CreateOrderUseCase:  createOrderUseCase,
		ListOrderesUseCase:  listOrderesUseCase,
		DeleteOrderUseCase:  deleteOrderUseCase,
		GetOrderByIDUseCase: getOrderByIDUseCase,
		UpdateOrderUseCase:  updateOrderUseCase,
	}
}

func (dh *OrderHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		Orderes := api.Group("/orders")
		{
			Orderes.POST("/", dh.CreateOrderHandler)
			Orderes.GET("/", dh.ListOrderesHandler)
			Orderes.DELETE("/:id", dh.DeleteOrderHandler)
			Orderes.GET("/:id", dh.GetOrderByIDHandler)
			Orderes.PUT("/:id", dh.UpdateOrderHandler)
		}
	}
}

func (dh *OrderHandlers) CreateOrderHandler(c *gin.Context) {
	var input usecase.CreateOrderInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := dh.CreateOrderUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (dh *OrderHandlers) ListOrderesHandler(c *gin.Context) {
	output, err := dh.ListOrderesUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dh *OrderHandlers) DeleteOrderHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.DeleteOrderInputDto{ID: id}
	output, err := dh.DeleteOrderUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dh *OrderHandlers) GetOrderByIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.GetOrderByIDInputDto{ID: id}
	output, err := dh.GetOrderByIDUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dh *OrderHandlers) UpdateOrderHandler(c *gin.Context) {
	id := c.Param("id")
	var input usecase.UpdateOrderInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id

	output, err := dh.UpdateOrderUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
