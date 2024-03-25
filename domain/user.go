package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"user_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Username  string    `json:"user_name" valid:"notnull"`
	Email     string    `json:"email" valid:"notnull"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func NewUser(user_name, email string) (*User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id.String(),
		Username: user_name,
		Email:    email,
	}, nil
}
