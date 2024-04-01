package usecase

import (
	"evaeats/user-service/internal/user/entity"
	"time"
)

type CreateUserInputDto struct {
	UserName     string `json:"user_name" valid:"notnull"`
	Email        string `json:"email" valid:"notnull"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserType     string `json:"user_type" validate:"eq=PRO|eq=CLIENT"`
	ProfileImage string `json:"profile_image"`
}

type CreateUserOutputDto struct {
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

type CreateUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewCreateUserUseCase(UserRepository entity.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: UserRepository}
}

func (u *CreateUserUseCase) Execute(input CreateUserInputDto) (*CreateUserOutputDto, error) {
	User := entity.NewUser(
		input.UserName,
		input.Email,
	)

	err := u.UserRepository.Create(User)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDto{
		ID:           User.ID,
		UserName:     User.UserName,
		Email:        User.Email,
		Password:     User.Password,
		Phone:        User.Phone,
		FirstName:    User.FirstName,
		LastName:     User.LastName,
		UserType:     User.UserType,
		ProfileImage: User.ProfileImage,
		CreatedAt:    User.CreatedAt,
		UpdatedAt:    User.UpdatedAt,
	}, nil
}
