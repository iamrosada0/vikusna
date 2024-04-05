package repository

import (
	"errors"
	"evaeats/user-service/internal/cheff/entity"
	userEntity "evaeats/user-service/internal/user/entity"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CheffRepositoryPostgres struct {
	DB *gorm.DB
}

func NewCheffRepositoryPostgres(db *gorm.DB) *CheffRepositoryPostgres {
	return &CheffRepositoryPostgres{DB: db}
}

func (r *CheffRepositoryPostgres) Create(cheff *entity.Cheff) error {
	// Check if the user exists
	var user userEntity.User
	err := r.DB.Where("ID = ?", cheff.UserId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Check if a chef with the same user ID already exists
	var existingChef entity.Cheff
	err = r.DB.Where("user_id = ?", cheff.UserId).First(&existingChef).Error
	if err == nil {
		return errors.New("chef with the same user ID already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// User exists and no chef with the same user ID, create the chef
	return r.DB.Create(cheff).Error
}

func (r *CheffRepositoryPostgres) FindAll() ([]*entity.Cheff, error) {
	var cheffs []*entity.Cheff
	if err := r.DB.Find(&cheffs).Error; err != nil {
		return nil, err
	}
	return cheffs, nil
}

func (r *CheffRepositoryPostgres) Update(Cheff *entity.Cheff) error {
	return r.DB.Save(Cheff).Error
}

func (r *CheffRepositoryPostgres) DeleteByID(id string) error {
	return r.DB.Where("id = ?", id).Delete(entity.Cheff{}).Error
}

func (r *CheffRepositoryPostgres) GetByID(id string) (*entity.Cheff, error) {
	var cheff entity.Cheff
	if err := r.DB.Where("id = ?", id).First(&cheff).Error; err != nil {
		return nil, err
	}
	return &cheff, nil
}
