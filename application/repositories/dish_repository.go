package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type DishRepository interface {
	Insert(chef_id, name, description string, price float64, availble bool) (*domain.Dish, error)
	Find(id string) (*domain.Dish, error)
}

type DishRepositoryDb struct {
	Db *gorm.DB
}

func (repo DishRepositoryDb) Insert(chef_id, name, description string, price float64, availble bool) (*domain.Dish, error) {
	// Generate a new Dish with a generated ID
	newDish, err := domain.NewDish(chef_id, name, description, price, availble)
	if err != nil {
		return nil, err
	}

	// Insert the Dish into the database
	if err := repo.Db.Create(newDish).Error; err != nil {
		return nil, err
	}

	return newDish, nil
}

func (repo DishRepositoryDb) Find(id string) (*domain.Dish, error) {
	var dish domain.Dish
	if err := repo.Db.First(&dish, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Dish with ID %s not found", id)
		}
		return nil, err
	}
	return &dish, nil
}
