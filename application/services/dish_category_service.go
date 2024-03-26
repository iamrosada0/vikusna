package services

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type DishCategoryService struct {
	DishCategoryRepository repositories.DishCategoryRepository
}

func NewDishCategoryService(dishCategoryRepo repositories.DishCategoryRepository) *DishCategoryService {
	return &DishCategoryService{
		DishCategoryRepository: dishCategoryRepo,
	}
}

func (s *DishCategoryService) CreateDishCategory(name string) (*domain.DishCategory, error) {
	// Validate input data, if necessary

	// Create the dish category in the database
	newDishCategory, err := s.DishCategoryRepository.Insert(name)
	if err != nil {
		return nil, err
	}

	return newDishCategory, nil
}

func (s *DishCategoryService) GetDishCategoryByID(id string) (*domain.DishCategory, error) {
	// Fetch the dish category by ID from the database
	dishCategory, err := s.DishCategoryRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return dishCategory, nil
}
