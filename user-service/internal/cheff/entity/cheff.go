package entity

import (
	"errors"
	"math/rand"
)

type CheffRepository interface {
	Create(cheff *Cheff) error
	FindAll() ([]*Cheff, error)
	Update(cheff *Cheff) error
	DeleteByID(id uint) error
	GetByID(id uint) (*Cheff, error)
}

type Cheff struct {
	ID    uint
	Name  string
	Email string
}

func NewCheff(name, email string) *Cheff {
	return &Cheff{
		ID:    uint(rand.Uint32()),
		Name:  name,
		Email: email,
	}
}

func (d *Cheff) Update(name, email string) {
	d.Name = name
	d.Email = email
}

type InMemoryCheffRepository struct {
	Cheffs map[string]*Cheff
}

func NewInMemoryCheffRepository() *InMemoryCheffRepository {
	return &InMemoryCheffRepository{
		Cheffs: make(map[string]*Cheff),
	}
}

func (r *InMemoryCheffRepository) DeleteByID(id string) error {
	if _, exists := r.Cheffs[id]; !exists {
		return errors.New("Cheff not found")
	}

	delete(r.Cheffs, id)
	return nil
}

func (r *InMemoryCheffRepository) FindAll() ([]*Cheff, error) {
	var allCheffs []*Cheff
	for _, cheff := range r.Cheffs {
		allCheffs = append(allCheffs, cheff)
	}
	return allCheffs, nil
}
