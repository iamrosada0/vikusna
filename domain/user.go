package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string    `json:"user_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	User_name    string    `json:"user_name" valid:"notnull"`
	Email        string    `json:"email" valid:"notnull"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	First_name   string    `json:"first_name"`
	Last_name    string    `json:"last_name"`
	User_type    string    `json:"user_type" validate:"eq=PRO|eq=CLIENT"`
	ProfileImage string    `json:"profile_image"`
	Profile      []Profile `gorm:"profile"`

	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func NewUser(user_name, email, password, phone, first_name, last_name, user_type, profile_image string) (*User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           id.String(),
		User_name:    user_name,
		Email:        email,
		Password:     password,
		Phone:        phone,
		First_name:   first_name,
		Last_name:    last_name,
		User_type:    user_type,
		ProfileImage: profile_image,
	}, nil
}
