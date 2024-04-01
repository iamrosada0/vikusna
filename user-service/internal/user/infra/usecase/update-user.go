package usecase

import (
	"evaeats/user-service/internal/user/entity"
	"time"
)

type UpdateUserInputDto struct {
	ID           string    `json:"user_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	UserName     string    `json:"user_name" valid:"notnull"`
	Email        string    `json:"email" valid:"notnull"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	UserType     string    `json:"user_type" validate:"eq=PRO|eq=CLIENT"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at" valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

type UpdateUserOutputDto struct {
	ID           string    `json:"user_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	UserName     string    `json:"user_name" valid:"notnull"`
	Email        string    `json:"email" valid:"notnull"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	UserType     string    `json:"user_type" validate:"eq=PRO|eq=CLIENT"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at" valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

type UpdateUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewUpdateUserUseCase(UserRepository entity.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserRepository: UserRepository}
}

func (u *UpdateUserUseCase) Execute(input UpdateUserInputDto) (*UpdateUserOutputDto, error) {
	existingUser, err := u.UserRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existingUser.Update(
		input.UserName,
		input.Email,
		input.Password,
		input.Phone,
		input.FirstName,
		input.LastName,
		input.UserType,
		input.ProfileImage,
	)

	err = u.UserRepository.Update(existingUser)
	if err != nil {
		return nil, err
	}

	return &UpdateUserOutputDto{
		ID:           existingUser.ID,
		UserName:     existingUser.UserName,
		Email:        existingUser.Email,
		Password:     existingUser.Password,
		Phone:        existingUser.Phone,
		FirstName:    existingUser.FirstName,
		LastName:     existingUser.LastName,
		UserType:     existingUser.UserType,
		ProfileImage: existingUser.ProfileImage,
		CreatedAt:    existingUser.CreatedAt,
		UpdatedAt:    existingUser.UpdatedAt,
	}, nil
}
