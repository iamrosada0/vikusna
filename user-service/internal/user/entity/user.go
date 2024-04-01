package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]*User, error)
	Update(user *User) error
	DeleteByID(id string) error
	GetByID(id string) (*User, error)
}

type User struct {
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

func NewUser(name, email string) *User {
	return &User{
		ID:       uuid.New().String(),
		UserName: name,
		Email:    email,
	}
}

func (u *User) Update(userName, email, password, phone, firstName, lastName, userType, profileImage string) {
	u.UserName = userName
	u.Email = email
	u.Password = password
	u.Phone = phone
	u.FirstName = firstName
	u.LastName = lastName
	u.UserType = userType
	u.ProfileImage = profileImage
	u.UpdatedAt = time.Now()
}

type InMemoryUserRepository struct {
	users map[string]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*User),
	}
}

func (r *InMemoryUserRepository) Create(user *User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) DeleteByID(id string) error {
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

func (r *InMemoryUserRepository) FindAll() ([]*User, error) {
	var allUsers []*User
	for _, user := range r.users {
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}

func (r *InMemoryUserRepository) GetByID(id string) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) Update(user *User) error {
	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}
