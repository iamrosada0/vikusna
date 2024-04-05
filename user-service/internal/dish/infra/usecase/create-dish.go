package usecase

import (
	"evaeats/user-service/internal/dish/entity"
	"fmt"
)

type CreateDishUseCase struct {
	DishRepository entity.DishRepository
}

func NewCreateDishUseCase(DishRepository entity.DishRepository) *CreateDishUseCase {
	return &CreateDishUseCase{DishRepository: DishRepository}
}
func (input CreateDishInputDto) Validate() error {
	if input.ChefID == "" {
		return fmt.Errorf("ChefID cannot be empty")
	}
	if input.Name == "" {
		return fmt.Errorf("Name cannot be empty")
	}
	if input.Description == "" {
		return fmt.Errorf("Description cannot be empty")
	}
	if input.Dish_image == "" {
		return fmt.Errorf("Dish_image cannot be empty")
	}
	if input.Price <= 0 {
		return fmt.Errorf("Price must be greater than zero")
	}
	// Add validation for CategoryID if required
	if input.CategoryID == "" {
		return fmt.Errorf("CategoryID cannot be empty")
	}
	return nil
}

func (u *CreateDishUseCase) Execute(input CreateDishInputDto) (*CreateDishOutputDto, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// Create a new Dish entity using input data
	newDish, err := entity.NewDish(
		input.ChefID,
		input.Name,
		input.Description,
		input.Dish_image,
		input.Price,
		input.Available,
		input.CategoryID, // Added category ID to create dish
	)
	if err != nil {
		return nil, err
	}

	// Call DishRepository to create the Dish
	err = u.DishRepository.Create(newDish)
	if err != nil {
		return nil, err
	}

	// Construct output DTO using created Dish
	output := &CreateDishOutputDto{
		ID:          newDish.ID,
		ChefID:      newDish.ChefID,
		Name:        newDish.Name,
		Description: newDish.Description,
		Dish_image:  newDish.DishImage,
		Price:       newDish.Price,
		Available:   newDish.Available,
	}

	return output, nil
}

// CreateDishInputDto defines the input data structure for creating a dish
type CreateDishInputDto struct {
	ChefID      string  `json:"chef_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Dish_image  string  `json:"dish_image"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
	CategoryID  string  `json:"category_id"` // Added category ID
}

// CreateDishOutputDto defines the output data structure for creating a dish
type CreateDishOutputDto struct {
	ID          string  `json:"dish_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ChefID      string  `json:"chef_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Dish_image  string  `json:"dish_image"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
}
