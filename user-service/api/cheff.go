package api

import (
	"evaeats/user-service/internal/cheff/infra/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheffHandlers struct {
	CreateCheffUseCase  *usecase.CreateCheffUseCase
	ListCheffsUseCase   *usecase.GetAllCheffsUseCase
	DeleteCheffUseCase  *usecase.DeleteCheffUseCase
	GetCheffByIDUseCase *usecase.GetCheffByIDUseCase
	UpdateCheffUseCase  *usecase.UpdateCheffUseCase
}

func NewCheffHandlers(
	createCheffUseCase *usecase.CreateCheffUseCase,
	listCheffsUseCase *usecase.GetAllCheffsUseCase,
	deleteCheffUseCase *usecase.DeleteCheffUseCase,
	getCheffByIDUseCase *usecase.GetCheffByIDUseCase,
	updateCheffUseCase *usecase.UpdateCheffUseCase,
) *CheffHandlers {
	return &CheffHandlers{
		CreateCheffUseCase:  createCheffUseCase,
		ListCheffsUseCase:   listCheffsUseCase,
		DeleteCheffUseCase:  deleteCheffUseCase,
		GetCheffByIDUseCase: getCheffByIDUseCase,
		UpdateCheffUseCase:  updateCheffUseCase,
	}
}

func (ch *CheffHandlers) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		cheffs := api.Group("/cheffs")
		{
			cheffs.POST("/", ch.CreateCheffHandler)
			cheffs.GET("/", ch.ListCheffsHandler)
			cheffs.DELETE("/", ch.DeleteCheffHandler)
			cheffs.GET("/:id", ch.GetCheffByIDHandler)
			cheffs.PUT("/", ch.UpdateCheffHandler)
		}
	}
}

func (ch *CheffHandlers) CreateCheffHandler(c *gin.Context) {
	var input usecase.CreateCheffInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := ch.CreateCheffUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}

func (ch *CheffHandlers) ListCheffsHandler(c *gin.Context) {
	output, err := ch.ListCheffsUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (ch *CheffHandlers) DeleteCheffHandler(c *gin.Context) {
	var input usecase.DeleteCheffInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := ch.DeleteCheffUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (ch *CheffHandlers) GetCheffByIDHandler(c *gin.Context) {
	id := c.Param("id")

	input := usecase.GetCheffByIDInputDto{ID: id}
	output, err := ch.GetCheffByIDUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (ch *CheffHandlers) UpdateCheffHandler(c *gin.Context) {
	var input usecase.UpdateCheffInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := ch.UpdateCheffUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
