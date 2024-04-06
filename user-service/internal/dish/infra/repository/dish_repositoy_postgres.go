package repository

import (
	"errors"
	"evaeats/user-service/internal/dish/entity"

	cheffEntity "evaeats/user-service/internal/cheff/entity"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DishRepositoryPostgres struct {
	DB *gorm.DB
}

func NewDishRepositoryPostgres(db *gorm.DB) *DishRepositoryPostgres {
	return &DishRepositoryPostgres{DB: db}
}

func (r *DishRepositoryPostgres) Create(Dish *entity.Dish) error {
	// Verificar se o chef_id existe na tabela de chefs
	var chef cheffEntity.Cheff
	if err := r.DB.Where("id = ?", Dish.ChefID).First(&chef).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("chef not found")
		}
		return err
	}

	// Verificar se o category_id existe na tabela de DishCategory
	var category entity.DishCategory
	if err := r.DB.Where("id = ?", Dish.CategoryID).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category not found")
		}
		return err
	}

	// Verificar se já existe um prato com o mesmo nome e descrição para o chef
	var existingDish entity.Dish
	if err := r.DB.Where("chef_id = ? AND name = ? AND description = ?", Dish.ChefID, Dish.Name, Dish.Description).First(&existingDish).Error; err == nil {
		return errors.New("a dish with the same name and description already exists for this chef")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Se todas as verificações passarem, criar o prato (dish)
	return r.DB.Create(Dish).Error
}

func (r *DishRepositoryPostgres) FindAll() ([]*entity.Dish, error) {
	var Dishs []*entity.Dish
	if err := r.DB.Preload("DishCategory").Find(&Dishs).Error; err != nil {
		return nil, err
	}
	return Dishs, nil
}

func (r *DishRepositoryPostgres) Update(Dish *entity.Dish) error {
	return r.DB.Save(Dish).Error
}

func (r *DishRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(entity.Dish{}).Error
}

func (r *DishRepositoryPostgres) GetByID(id string) (*entity.Dish, error) {
	var Dish entity.Dish
	if err := r.DB.Where("id = ?", id).First(&Dish).Error; err != nil {
		return nil, err
	}
	return &Dish, nil
}

func (r *DishRepositoryPostgres) FindByCategoryName(categoryName string) ([]*entity.Dish, error) {
	var dishesInCategory []*entity.Dish

	// Query dishes where category name matches
	if err := r.DB.
		Joins("JOIN dish_categories ON dish_categories.id = dishes.category_id").
		Where("dish_categories.name = ?", categoryName).
		Find(&dishesInCategory).Error; err != nil {
		return nil, err
	}

	return dishesInCategory, nil
}
