package usecase

import "evaeats/user-service/internal/dish/entity"

type CreateDishInputDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateDishOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateDishUseCase struct {
	DishRepository entity.DishRepository
}

func NewCreateDishUseCase(DishRepository entity.DishRepository) *CreateDishUseCase {
	return &CreateDishUseCase{DishRepository: DishRepository}
}

func (u *CreateDishUseCase) Execute(input CreateDishInputDto) (*CreateDishOutputDto, error) {
	Dish := entity.NewDish(
		input.Name,
		input.Email,
	)

	err := u.DishRepository.Create(Dish)
	if err != nil {
		return nil, err
	}

	return &CreateDishOutputDto{
		ID:    Dish.ID,
		Name:  Dish.Name,
		Email: Dish.Email,
	}, nil
}
