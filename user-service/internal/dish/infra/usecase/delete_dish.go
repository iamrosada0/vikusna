package usecase

import "evaeats/user-service/internal/dish/entity"

type DeleteDishInputDto struct {
	ID uint `json:"id"`
}

type DeleteDishOutputDto struct {
	ID uint `json:"id"`
}

type DeleteDishUseCase struct {
	DishRepository entity.DishRepository
}

func NewDeleteDishUseCase(DishRepository entity.DishRepository) *DeleteDishUseCase {
	return &DeleteDishUseCase{DishRepository: DishRepository}
}

func (u *DeleteDishUseCase) Execute(input DeleteDishInputDto) (*DeleteDishOutputDto, error) {
	err := u.DishRepository.DeleteByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteDishOutputDto{ID: input.ID}, nil
}
