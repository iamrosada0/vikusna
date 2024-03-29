package repository

import (
	"evaeats/user-service/internal/dish/entity"

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
	return r.DB.Create(Dish).Error
}

func (r *DishRepositoryPostgres) FindAll() ([]*entity.Dish, error) {
	var Dishs []*entity.Dish
	if err := r.DB.Find(&Dishs).Error; err != nil {
		return nil, err
	}
	return Dishs, nil
}

func (r *DishRepositoryPostgres) Update(Dish *entity.Dish) error {
	return r.DB.Save(Dish).Error
}

func (r *DishRepositoryPostgres) DeleteByID(id uint) error {
	return r.DB.Where("id = ?", id).Delete(entity.Dish{}).Error
}

func (r *DishRepositoryPostgres) GetByID(id uint) (*entity.Dish, error) {
	var Dish entity.Dish
	if err := r.DB.Where("id = ?", id).First(&Dish).Error; err != nil {
		return nil, err
	}
	return &Dish, nil
}
