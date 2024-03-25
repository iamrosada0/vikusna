package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type DishCategoryRepository interface {
	Insert(name string) (*domain.DishCategory, error)
	Find(id string) (*domain.DishCategory, error)
}

type DishCategoryRepositoryDb struct {
	Db *gorm.DB
}

func (repo DishCategoryRepositoryDb) Insert(name string) (*domain.DishCategory, error) {
	// Generate a new DishCategory with a generated ID
	newDishCategory, err := domain.NewDishCategory(name)
	if err != nil {
		return nil, err
	}

	// Insert the DishCategory into the database
	if err := repo.Db.Create(newDishCategory).Error; err != nil {
		return nil, err
	}

	return newDishCategory, nil
}

func (repo DishCategoryRepositoryDb) Find(id string) (*domain.DishCategory, error) {
	var dishCategory domain.DishCategory
	if err := repo.Db.First(&dishCategory, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("dish category with ID %s not found", id)
		}
		return nil, err
	}
	return &dishCategory, nil
}
