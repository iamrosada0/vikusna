package domain

import "github.com/google/uuid"

type Dish struct {
	ID          string  `json:"dish_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ChefID      string  `json:"chef_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Dish_image  string  `json:"dish_image"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
}

func NewDish(chefID, name, description, dish_image string, price float64, available bool) (*Dish, error) {
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
		Dish_image:  dish_image,
		Available:   available,
	}, nil
}
