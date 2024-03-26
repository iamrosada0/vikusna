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

func (s *UserService) CreateUser(userName, email string) (*domain.User, error) {
	return s.UserRepository.Insert(userName, email)
}

func (s *UserService) GetUserByID(userID string) (*domain.User, error) {
	return s.UserRepository.Find(userID)
}
