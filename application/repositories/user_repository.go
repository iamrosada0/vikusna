package repositories

import (
	"evaeats/domain"
	"evaeats/helper"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var profile_image = ""

type UserRepository interface {
	Insert(user_name, email, password, phone, first_name, last_name, user_type, profile_image string) (*domain.User, error)
	Find(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByPhone(phone string) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user_name, email, password, phone, first_name, last_name, user_type, profile_image string) (*domain.User, error) {
	var user domain.User
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	usernameEmailStrip := strings.Split(email, "@")[0]

	pro_type := ""

	if user.User_type == "CLIENT" {
		pro_type = ""
	} else {
		pro_type = "CHEF"
	}

	newUser := &domain.User{
		ID:           id.String(),
		User_name:    user_name,
		Email:        email,
		Password:     helper.HashPassword(user.Password),
		Phone:        phone,
		First_name:   first_name,
		Last_name:    last_name,
		User_type:    user_type,
		ProfileImage: profile_image,
		Profile: []domain.Profile{{
			First_name: first_name,
			Last_name:  last_name,
			User_type:  user_type,
			User_name:  usernameEmailStrip,
			Pro_type:   pro_type,
		}},
	}

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

func (repo UserRepositoryDb) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := repo.Db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Not found
		}
		return nil, err
	}
	return &user, nil
}

func (repo UserRepositoryDb) FindByPhone(phone string) (*domain.User, error) {
	var user domain.User
	if err := repo.Db.Where("phone = ?", phone).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Not found
		}
		return nil, err
	}
	return &user, nil
}
