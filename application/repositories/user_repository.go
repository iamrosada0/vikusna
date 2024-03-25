package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user_name, email string) (*domain.User, error)
	Find(id string) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user_name, email string) (*domain.User, error) {
	// Generate a new user with a generated ID
	newUser, err := domain.NewUser(user_name, email)
	if err != nil {
		return nil, err
	}

	// Insert the user into the database
	if err := repo.Db.Create(newUser).Error; err != nil {
		return nil, err
	}

	return newUser, nil
}

func (repo UserRepositoryDb) Find(id string) (*domain.User, error) {
	var user domain.User
	if err := repo.Db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User with ID %s not found", id)
		}
		return nil, err
	}
	return &user, nil
}
