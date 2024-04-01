package usecase

import (
	"evaeats/user-service/internal/dish/entity"
)

type CreateDishUseCase struct {
	DishRepository entity.DishRepository
}

func NewCreateDishUseCase(DishRepository entity.DishRepository) *CreateDishUseCase {
	return &CreateDishUseCase{DishRepository: DishRepository}
}

func (u *CreateDishUseCase) Execute(input CreateDishInputDto) (*CreateDishOutputDto, error) {
	// Create a new Dish entity using input data
	newDish, err := entity.NewDish(
		input.ChefID,
		input.Name,
		input.Description,
		input.Dish_image,
		input.Price,
		input.Available,
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
		Dish_image:  newDish.Dish_image,
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
