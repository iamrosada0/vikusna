package repository

import (
	"errors"
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
	// Verifica se já existe uma categoria com o mesmo nome
	var existingCategory entity.DishCategory
	result := r.DB.Where("name = ?", category.Name).First(&existingCategory)
	if result.Error == nil {
		// Categoria com o mesmo nome já existe, retornar erro com o nome existente
		return errors.New("categoria com o nome '" + existingCategory.Name + "' já existe")
	}

	// Cria a categoria se não houver erro de consulta
	if err := r.DB.Create(category).Error; err != nil {
		return err
	}

	return nil
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
