package usecase

import "evaeats/user-service/internal/dish/entity"

type GetDishByIDInputDto struct {
	ID string `json:"id"`
}

type GetDishByIDOutputDto struct {
	ID          string  `json:"dish_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ChefID      string  `json:"chef_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Dish_image  string  `json:"dish_image"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
}

type GetDishByIDUseCase struct {
	DishRepository entity.DishRepository
}

func NewGetDishByIDUseCase(DishRepository entity.DishRepository) *GetDishByIDUseCase {
	return &GetDishByIDUseCase{DishRepository: DishRepository}
}

func (u *GetDishByIDUseCase) Execute(input GetDishByIDInputDto) (*GetDishByIDOutputDto, error) {
	Dish, err := u.DishRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetDishByIDOutputDto{
		ID:          Dish.ID,
		ChefID:      Dish.ChefID,
		Name:        Dish.Name,
		Description: Dish.Description,
		Dish_image:  Dish.DishImage,
		Price:       Dish.Price,
		Available:   Dish.Available,
	}, nil
}
