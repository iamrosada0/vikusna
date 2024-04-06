package entity

import (
	"errors"

	"github.com/google/uuid"
)

type DishRepository interface {
	Create(Dish *Dish) error
	FindAll() ([]*Dish, error)
	Update(Dish *Dish) error
	DeleteByID(id string) error
	GetByID(id string) (*Dish, error)
	FindByCategoryName(categoryName string) ([]*Dish, error) // Nova função adicionada

}

type Dish struct {
	ID           string       `json:"dish_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	ChefID       string       `json:"chef_id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	DishImage    string       `json:"dish_image"`
	Price        float64      `json:"price"`
	Available    bool         `json:"available"`
	CategoryID   string       `json:"category_id"`
	DishCategory DishCategory `gorm:"foreignkey:CategoryID"`
}

func NewDish(chefID, name, description, dish_image string, price float64, available bool, categoryID string) (*Dish, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Dish{
		ID:          id.String(),
		ChefID:      chefID,
		Name:        name,
		Description: description,
		DishImage:   dish_image,
		Price:       price,
		Available:   available,
		CategoryID:  categoryID,
	}, nil
}

func (d *Dish) Update(name, description, dishImage string, price float64, available bool) {
	d.Name = name
	d.Description = description
	d.DishImage = dishImage
	d.Price = price
	d.Available = available
}

type InMemoryDishRepository struct {
	Dishes map[string]*Dish
}

func NewInMemoryDishRepository() *InMemoryDishRepository {
	return &InMemoryDishRepository{
		Dishes: make(map[string]*Dish),
	}
}

func (r *InMemoryDishRepository) Create(Dish *Dish) error {
	// Implement the logic to create a dish
	// For an in-memory repository, you would add the dish to the map
	r.Dishes[Dish.ID] = Dish
	return nil
}

func (r *InMemoryDishRepository) DeleteByID(id string) error {
	if _, exists := r.Dishes[id]; !exists {
		return errors.New("Dish not found")
	}

	delete(r.Dishes, id)
	return nil
}

func (r *InMemoryDishRepository) FindAll() ([]*Dish, error) {
	var allDishes []*Dish
	for _, dish := range r.Dishes {
		allDishes = append(allDishes, dish)
	}
	return allDishes, nil
}

func (r *InMemoryDishRepository) Update(Dish *Dish) error {
	// Implement the logic to update a dish
	// For an in-memory repository, you would update the dish in the map
	r.Dishes[Dish.ID] = Dish
	return nil
}

func (r *InMemoryDishRepository) GetByID(id string) (*Dish, error) {
	dish, exists := r.Dishes[id]
	if !exists {
		return nil, errors.New("Dish not found")
	}
	return dish, nil
}
