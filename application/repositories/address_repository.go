package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type AddressRepository interface {
	Insert(customer_id, street, city, state, postal_code string) (*domain.Address, error)
	Find(id string) (*domain.Address, error)
}

type AddressRepositoryDb struct {
	Db *gorm.DB
}

func (repo AddressRepositoryDb) Insert(customer_id, street, city, state, postal_code string) (*domain.Address, error) {
	// Generate a new Address with a generated ID
	newAddress, err := domain.NewAddress(customer_id, street, city, state, postal_code)
	if err != nil {
		return nil, err
	}

	// Insert the Address into the database
	if err := repo.Db.Create(newAddress).Error; err != nil {
		return nil, err
	}

	return newAddress, nil
}

func (repo AddressRepositoryDb) Find(id string) (*domain.Address, error) {
	var address domain.Address
	if err := repo.Db.First(&address, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("address with ID %s not found", id)
		}
		return nil, err
	}
	return &address, nil
}
