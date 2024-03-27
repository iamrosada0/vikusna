package services

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type DishService struct {
	DishRepository repositories.DishRepository
}

func NewDishService(dishRepo repositories.DishRepository) *DishService {
	return &DishService{
		DishRepository: dishRepo,
	}
}

func (s *DishService) CreateDish(chefID, name, description, dish_image string, price float64, available bool) (*domain.Dish, error) {
	// Validate input data, if necessary

	// Create the dish in the database
	newDish, err := s.DishRepository.Insert(chefID, name, description, dish_image, price, available)
	if err != nil {
		return nil, err
	}

	return newDish, nil
}

func (s *DishService) GetDishByID(id string) (*domain.Dish, error) {
	// Fetch the dish by ID from the database
	dish, err := s.DishRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return dish, nil
}
