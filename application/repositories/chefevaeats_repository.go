package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type ChefEvaEatsRepository interface {
	Insert(chefEvaEatsImage, chefEvaEatsName, phoneNumber, address, locationID, profileID, registrationStatus string) (*domain.ChefEvaEats, error)
	Find(id string) (*domain.ChefEvaEats, error)
	Update(id string, chefEvaEats *domain.ChefEvaEats) (*domain.ChefEvaEats, error)
}

type ChefEvaEatsRepositoryDb struct {
	Db *gorm.DB
}

func NewChefEvaEatsRepository(Db *gorm.DB) *ChefEvaEatsRepositoryDb {
	return &ChefEvaEatsRepositoryDb{Db: Db}
}

func (repo *ChefEvaEatsRepositoryDb) Insert(chefEvaEatsImage, chefEvaEatsName, phoneNumber, address, locationID, profileID, registrationStatus string) (*domain.ChefEvaEats, error) {
	var dbChef domain.ChefEvaEats
	repo.Db.Where("chef_evaeats_name = ?", chefEvaEatsName).First(&dbChef)
	if dbChef.ID != "" {
		return nil, fmt.Errorf("ChefEvaEats with the name %s already exists", chefEvaEatsName)
	}

	dbChef.Registration_status = "PENDING"
	if err := repo.Db.Create(&dbChef).Error; err != nil {
		return nil, err
	}

	chef := &domain.ChefEvaEats{
		ChefEvaEats_image:   chefEvaEatsImage,
		ChefEvaEats_name:    chefEvaEatsName,
		Phone_number:        phoneNumber,
		Address:             address,
		LocationID:          locationID,
		ProfileID:           profileID,
		Registration_status: registrationStatus,
	}
	if err := repo.Db.Create(chef).Error; err != nil {
		return nil, err
	}
	return chef, nil
}

func (repo *ChefEvaEatsRepositoryDb) Find(id string) (*domain.ChefEvaEats, error) {
	var chef domain.ChefEvaEats
	if err := repo.Db.First(&chef, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &chef, nil
}

func (repo *ChefEvaEatsRepositoryDb) Update(id string, chefEvaEats *domain.ChefEvaEats) (*domain.ChefEvaEats, error) {
	var dbChefEvaEats domain.ChefEvaEats

	if err := repo.Db.First(&dbChefEvaEats, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := repo.Db.Model(&dbChefEvaEats).Updates(chefEvaEats).Error; err != nil {
		return nil, err
	}

	return chefEvaEats, nil
}
