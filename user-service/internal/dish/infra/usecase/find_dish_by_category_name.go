package usecase

import "evaeats/user-service/internal/dish/entity"

type FindDishesByCategoryNameInputDto struct {
	CategoryName string `json:"category_name"`
}

type FindDishesByCategoryNameOutputDto struct {
	Dishes []*entity.Dish `json:"dishes"`
}

type FindDishesByCategoryNameUseCase struct {
	DishRepository entity.DishRepository
}

func NewFindDishesByCategoryNameUseCase(dishRepository entity.DishRepository) *FindDishesByCategoryNameUseCase {
	return &FindDishesByCategoryNameUseCase{DishRepository: dishRepository}
}

func (u *FindDishesByCategoryNameUseCase) Execute(input FindDishesByCategoryNameInputDto) (*FindDishesByCategoryNameOutputDto, error) {
	dishes, err := u.DishRepository.FindByCategoryName(input.CategoryName)
	if err != nil {
		return nil, err
	}

	return &FindDishesByCategoryNameOutputDto{
		Dishes: dishes,
	}, nil
}
