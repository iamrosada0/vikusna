package usecase

import "evaeats/user-service/internal/dish/entity"

type UpdateDishInputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateDishOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
		input.Email,
	)

	err = u.DishRepository.Update(existingDish)
	if err != nil {
		return nil, err
	}

	return &UpdateDishOutputDto{
		ID:    existingDish.ID,
		Name:  existingDish.Name,
		Email: existingDish.Email,
	}, nil
}
