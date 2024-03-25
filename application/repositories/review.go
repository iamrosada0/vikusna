package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	Insert(user_id, order_id, dish_id, comment string, rating int) (*domain.Review, error)
	Find(id string) (*domain.Review, error)
}

type ReviewRepositoryDb struct {
	Db *gorm.DB
}

func (repo ReviewRepositoryDb) Insert(user_id, order_id, dish_id, comment string, rating int) (*domain.Review, error) {
	// Generate a new Review with a generated ID
	newReview, err := domain.NewReview(user_id, order_id, dish_id, comment, rating)
	if err != nil {
		return nil, err
	}

	// Insert the Review into the database
	if err := repo.Db.Create(newReview).Error; err != nil {
		return nil, err
	}

	return newReview, nil
}

func (repo ReviewRepositoryDb) Find(id string) (*domain.Review, error) {
	var review domain.Review
	if err := repo.Db.First(&review, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("review with ID %s not found", id)
		}
		return nil, err
	}
	return &review, nil
}
