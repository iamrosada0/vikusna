package services

import (
	"evaeats/domain"
	"evaeats/user-service/internal/review/repository"
)

type ReviewService struct {
	ReviewRepository repository.ReviewRepository
}

func NewReviewService(reviewRepo repository.ReviewRepository) *ReviewService {
	return &ReviewService{
		ReviewRepository: reviewRepo,
	}
}

func (s *ReviewService) CreateReview(userID, orderID, dishID, comment string, rating int) (*domain.Review, error) {
	// Validate input data, if necessary

	// Create the review in the database
	newReview, err := s.ReviewRepository.Insert(userID, orderID, dishID, comment, rating)
	if err != nil {
		return nil, err
	}

	return newReview, nil
}

func (s *ReviewService) GetReviewByID(id string) (*domain.Review, error) {
	// Fetch the review by ID from the database
	review, err := s.ReviewRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return review, nil
}
