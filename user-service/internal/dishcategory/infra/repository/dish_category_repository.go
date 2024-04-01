package repository

import (
	"evaeats/user-service/internal/dish/entity"

	"gorm.io/gorm"
)

type DishCategoryRepositoryPostgres struct {
	DB *gorm.DB
}

func NewDishCategoryRepositoryPostgres(db *gorm.DB) *DishCategoryRepositoryPostgres {
	return &DishCategoryRepositoryPostgres{DB: db}
}

func (r *DishCategoryRepositoryPostgres) Create(category *entity.DishCategory) error {
	return r.DB.Create(category).Error
}

func (r *DishCategoryRepositoryPostgres) FindAll() ([]*entity.DishCategory, error) {
	var categories []*entity.DishCategory
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *DishCategoryRepositoryPostgres) Update(category *entity.DishCategory) error {
	return r.DB.Save(category).Error
}

func (r *DishCategoryRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(&entity.DishCategory{}).Error
}

func (r *DishCategoryRepositoryPostgres) GetByID(id string) (*entity.DishCategory, error) {
	var category entity.DishCategory
	if err := r.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
