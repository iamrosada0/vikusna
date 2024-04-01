package api

import (
	"evaeats/user-service/internal/dishcategory/infra/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DishCategoryHandlers struct {
	CreateCategoryUseCase  *usecase.CreateDishCategoryUseCase
	ListCategoriesUseCase  *usecase.GetAllDishCategoriesUseCase
	DeleteCategoryUseCase  *usecase.DeleteDishCategoryUseCase
	GetCategoryByIDUseCase *usecase.GetDishCategoryByIDUseCase
	UpdateCategoryUseCase  *usecase.UpdateDishCategoryUseCase
}

func NewDishCategoryHandlers(
	createCategoryUseCase *usecase.CreateDishCategoryUseCase,
	listCategoriesUseCase *usecase.GetAllDishCategoriesUseCase,
	deleteCategoryUseCase *usecase.DeleteDishCategoryUseCase,
	getCategoryByIDUseCase *usecase.GetDishCategoryByIDUseCase,
	updateCategoryUseCase *usecase.UpdateDishCategoryUseCase,
) *DishCategoryHandlers {
	return &DishCategoryHandlers{
		CreateCategoryUseCase:  createCategoryUseCase,
		ListCategoriesUseCase:  listCategoriesUseCase,
		DeleteCategoryUseCase:  deleteCategoryUseCase,
		GetCategoryByIDUseCase: getCategoryByIDUseCase,
		UpdateCategoryUseCase:  updateCategoryUseCase,
	}
}

func (dch *DishCategoryHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		categories := api.Group("/dish-categories")
		{
			categories.POST("/", dch.CreateCategoryHandler)
			categories.GET("/", dch.ListCategoriesHandler)
			categories.DELETE("/:id", dch.DeleteCategoryHandler)
			categories.GET("/:id", dch.GetCategoryByIDHandler)
			categories.PUT("/:id", dch.UpdateCategoryHandler)
		}
	}
}

func (dch *DishCategoryHandlers) CreateCategoryHandler(c *gin.Context) {
	var input usecase.CreateDishCategoryInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := dch.CreateCategoryUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (dch *DishCategoryHandlers) ListCategoriesHandler(c *gin.Context) {
	output, err := dch.ListCategoriesUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dch *DishCategoryHandlers) DeleteCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	err := dch.DeleteCategoryUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (dch *DishCategoryHandlers) GetCategoryByIDHandler(c *gin.Context) {
	id := c.Param("id")

	output, err := dch.GetCategoryByIDUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (dch *DishCategoryHandlers) UpdateCategoryHandler(c *gin.Context) {
	id := c.Param("id")
	var input usecase.UpdateDishCategoryInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id

	output, err := dch.UpdateCategoryUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
