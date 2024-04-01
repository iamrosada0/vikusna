package entity

import (
	"errors"

	"github.com/google/uuid"
)

type DishCategoryRepository interface {
	Create(category *DishCategory) error
	FindAll() ([]*DishCategory, error)
	Update(category *DishCategory) error
	DeleteByID(id string) error
	GetByID(id string) (*DishCategory, error)
}

type DishCategory struct {
	ID   string `json:"category_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name string `json:"name"`
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

type InMemoryDishCategoryRepository struct {
	Categories map[string]*DishCategory
}

func NewInMemoryDishCategoryRepository() *InMemoryDishCategoryRepository {
	return &InMemoryDishCategoryRepository{
		Categories: make(map[string]*DishCategory),
	}
}

func (r *InMemoryDishCategoryRepository) Create(category *DishCategory) error {
	if _, exists := r.Categories[category.ID]; exists {
		return errors.New("category already exists")
	}

	r.Categories[category.ID] = category
	return nil
}

func (r *InMemoryDishCategoryRepository) FindAll() ([]*DishCategory, error) {
	var allCategories []*DishCategory
	for _, category := range r.Categories {
		allCategories = append(allCategories, category)
	}
	return allCategories, nil
}

func (r *InMemoryDishCategoryRepository) Update(category *DishCategory) error {
	if _, exists := r.Categories[category.ID]; !exists {
		return errors.New("category not found")
	}

	r.Categories[category.ID] = category
	return nil
}

func (r *InMemoryDishCategoryRepository) DeleteByID(id string) error {
	if _, exists := r.Categories[id]; !exists {
		return errors.New("category not found")
	}

	delete(r.Categories, id)
	return nil
}

func (r *InMemoryDishCategoryRepository) GetByID(id string) (*DishCategory, error) {
	if category, exists := r.Categories[id]; exists {
		return category, nil
	}

	return nil, errors.New("category not found")
}
