package usecase

import "evaeats/user-service/internal/dish/entity"

type UpdateDishInputDto struct {
	ID          string  `json:"dish_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ChefID      string  `json:"chef_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Dish_image  string  `json:"dish_image"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
}

type UpdateDishOutputDto struct {
	ID          string  `json:"dish_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ChefID      string  `json:"chef_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Dish_image  string  `json:"dish_image"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
}

type UpdateDishUseCase struct {
	DishRepository entity.DishRepository
}

func NewUpdateDishUseCase(DishRepository entity.DishRepository) *UpdateDishUseCase {
	return &UpdateDishUseCase{DishRepository: DishRepository}
}

func (u *UpdateDishUseCase) Execute(input UpdateDishInputDto) (*UpdateDishOutputDto, error) {
	existingDish, err := u.DishRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existingDish.Update(
		input.Name,
		input.Description,
		input.Dish_image,
		input.Price,
		input.Available,
	)

	err = u.DishRepository.Update(existingDish)
	if err != nil {
		return nil, err
	}

	return &UpdateDishOutputDto{
		ID:          existingDish.ID,
		ChefID:      existingDish.ChefID,
		Name:        existingDish.Name,
		Description: existingDish.Description,
		Dish_image:  existingDish.Dish_image,
		Price:       existingDish.Price,
		Available:   existingDish.Available,
	}, nil
}
