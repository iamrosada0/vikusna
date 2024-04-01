package api

import (
	"evaeats/user-service/internal/dish/infra/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DishHandlers struct {
	CreateDishUseCase  *usecase.CreateDishUseCase
	ListDishesUseCase  *usecase.GetAllDishsUseCase
	DeleteDishUseCase  *usecase.DeleteDishUseCase
	GetDishByIDUseCase *usecase.GetDishByIDUseCase
	UpdateDishUseCase  *usecase.UpdateDishUseCase
}

func NewDishHandlers(
	createDishUseCase *usecase.CreateDishUseCase,
	listDishesUseCase *usecase.GetAllDishsUseCase,
	deleteDishUseCase *usecase.DeleteDishUseCase,
	getDishByIDUseCase *usecase.GetDishByIDUseCase,
	updateDishUseCase *usecase.UpdateDishUseCase,
) *DishHandlers {
	return &DishHandlers{
		CreateDishUseCase:  createDishUseCase,
		ListDishesUseCase:  listDishesUseCase,
		DeleteDishUseCase:  deleteDishUseCase,
		GetDishByIDUseCase: getDishByIDUseCase,
		UpdateDishUseCase:  updateDishUseCase,
	}
}

func (dh *DishHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		dishes := api.Group("/dishes")
		{
			dishes.POST("/", dh.CreateDishHandler)
			dishes.GET("/", dh.ListDishesHandler)
			dishes.DELETE("/:id", dh.DeleteDishHandler)
			dishes.GET("/:id", dh.GetDishByIDHandler)
			dishes.PUT("/:id", dh.UpdateDishHandler)
		}
	}
}

func (dh *DishHandlers) CreateDishHandler(c *gin.Context) {
	var input usecase.CreateDishInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := dh.CreateDishUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (dh *DishHandlers) ListDishesHandler(c *gin.Context) {
	output, err := dh.ListDishesUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dh *DishHandlers) DeleteDishHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.DeleteDishInputDto{ID: id}
	output, err := dh.DeleteDishUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dh *DishHandlers) GetDishByIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.GetDishByIDInputDto{ID: id}
	output, err := dh.GetDishByIDUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dh *DishHandlers) UpdateDishHandler(c *gin.Context) {
	id := c.Param("id")
	var input usecase.UpdateDishInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id

	output, err := dh.UpdateDishUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
