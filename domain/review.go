package domain

import "github.com/google/uuid"

// Review representa a avaliação de um cliente sobre um pedido ou prato específico
type Review struct {
	ID      string `json:"review_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	UserID  string
	OrderID string
	DishID  string
	Rating  int // de 1 a 5, por exemplo
	Comment string
}

func NewReview(userID, orderID, dishID, comment string, rating int) (*Review, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Review{
		ID:      id.String(),
		UserID:  userID,
		OrderID: orderID,
		DishID:  dishID,
		Rating:  rating,
		Comment: comment,
	}, nil
}
