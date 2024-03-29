package usecase

import "evaeats/user-service/internal/dish/entity"

type GetAllDishsOutputDto struct {
	Dishs []*entity.Dish `json:"Dishs"`
}

type GetAllDishsUseCase struct {
	DishRepository entity.DishRepository
}

func NewGetAllDishsUseCase(DishRepository entity.DishRepository) *GetAllDishsUseCase {
	return &GetAllDishsUseCase{DishRepository: DishRepository}
}

func (u *GetAllDishsUseCase) Execute() (*GetAllDishsOutputDto, error) {
	Dishs, err := u.DishRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &GetAllDishsOutputDto{Dishs: Dishs}, nil
}
