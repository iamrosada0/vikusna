package domain

import "github.com/google/uuid"

type Dish struct {
	ID          string `json:"dish_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ChefID      string
	Name        string
	Description string
	Price       float64
	Available   bool
}

func NewDish(chefID, name, description string, price float64, available bool) (*Dish, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Dish{
		ID:          id.String(),
		ChefID:      chefID,
		Name:        name,
		Description: description,
		Price:       price,
		Available:   available,
	}, nil
}
