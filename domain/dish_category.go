package domain

import "github.com/google/uuid"

// Categoria de Prato representa uma categoria na qual os pratos podem ser agrupados
type DishCategory struct {
	ID   string `json:"dish_category_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name string
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