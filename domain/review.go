package domain

import "github.com/google/uuid"

// Review representa a avaliação de um cliente sobre um pedido ou prato específico
type Review struct {
	ID      string `json:"review_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	UserID  string `json:"user_id"`
	OrderID string `json:"order_id"`
	DishID  string `json:"dish_id"`
	Rating  int    `json:"rating"` // de 1 a 5, por exemplo
	Comment string `json:"comment"`
}

// NewReview cria e retorna uma nova avaliação com um ID gerado automaticamente
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
