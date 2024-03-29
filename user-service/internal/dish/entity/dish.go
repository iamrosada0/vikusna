package entity

import (
	"errors"
	"math/rand"
)

type DishRepository interface {
	Create(Dish *Dish) error
	FindAll() ([]*Dish, error)
	Update(Dish *Dish) error
	DeleteByID(id uint) error
	GetByID(id uint) (*Dish, error)
}

type Dish struct {
	ID    uint
	Name  string
	Email string
}

func NewDish(name, email string) *Dish {
	return &Dish{
		ID:    uint(rand.Uint32()),
		Name:  name,
		Email: email,
	}
}

func (d *Dish) Update(name, email string) {
	d.Name = name
	d.Email = email
}

type InMemoryDishRepository struct {
	Dishs map[string]*Dish
}

func NewInMemoryDishRepository() *InMemoryDishRepository {
	return &InMemoryDishRepository{
		Dishs: make(map[string]*Dish),
	}
}

func (r *InMemoryDishRepository) DeleteByID(id string) error {
	if _, exists := r.Dishs[id]; !exists {
		return errors.New("Dish not found")
	}

	delete(r.Dishs, id)
	return nil
}

func (r *InMemoryDishRepository) FindAll() ([]*Dish, error) {
	var allDishs []*Dish
	for _, Dish := range r.Dishs {
		allDishs = append(allDishs, Dish)
	}
	return allDishs, nil
}
