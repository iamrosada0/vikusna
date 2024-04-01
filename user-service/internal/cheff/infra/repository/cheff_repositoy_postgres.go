package repository

import (
	"evaeats/user-service/internal/cheff/entity"

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
