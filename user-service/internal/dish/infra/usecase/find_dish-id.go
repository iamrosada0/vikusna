package usecase

import "evaeats/user-service/internal/dish/entity"

type GetDishByIDInputDto struct {
	ID uint `json:"id"`
}

type GetDishByIDOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
		ID:    Dish.ID,
		Name:  Dish.Name,
		Email: Dish.Email,
	}, nil
}
