package services

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) CreateUser(user_name, email, password, phone, first_name, last_name, user_type, profile_image string) (*domain.User, error) {

	return s.UserRepository.Insert(user_name, email, password, phone, first_name, last_name, user_type, profile_image)
}

func (s *UserService) GetUserByID(userID string) (*domain.User, error) {
	return s.UserRepository.Find(userID)
}
