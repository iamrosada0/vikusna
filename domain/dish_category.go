package domain

import "github.com/google/uuid"

// Categoria de Prato representa uma categoria na qual os pratos podem ser agrupados
type DishCategory struct {
	ID   string `json:"dish_category_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name string `json:"name"`

	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Food_image   string  `json:"food_image"`
	RestaurantID int     `json:"restarant_id"`
	Status       bool    `json:"status"`
}

func NewDishCategory(name string) (*DishCategory, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &DishCategory{
		ID:   id.String(),
		Name: name,
	}, nil
}
