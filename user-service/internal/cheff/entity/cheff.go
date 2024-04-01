package entity

import (
	"errors"

	"github.com/google/uuid"
)

type CheffRepository interface {
	Create(cheff *Cheff) error
	FindAll() ([]*Cheff, error)
	Update(cheff *Cheff) error
	DeleteByID(id string) error
	GetByID(id string) (*Cheff, error)
}

type Cheff struct {
	ID                  string `json:"cheff_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Cheff_image         string `json:"cheff_image"`
	Cheff_name          string `json:"cheff_name"`
	Phone_number        string `json:"phone_number"`
	Address             string `json:"address"`
	LocationID          string `json:"location"`
	Registration_status string `json:"registration_status" validate:"eq=PENDING|eq=ACCEPTED|eq=REJECTED"`
	LocationPreference  string `json:"location_preference"`
	CookingExperience   string `json:"cooking_experience"`
	Specialties         string `json:"specialties"`
	Certifications      string `json:"certifications"`
}

func NewCheff(
	cheff_image, cheff_name, phone_number, address, locationId,
	locationPreference, cookingExperience, specialties, certifications string,
) *Cheff {
	return &Cheff{
		ID:                  uuid.New().String(),
		Cheff_image:         cheff_image,
		Cheff_name:          cheff_name,
		Phone_number:        phone_number,
		Address:             address,
		LocationID:          locationId,
		Registration_status: "PENDING",
		LocationPreference:  locationPreference,
		CookingExperience:   cookingExperience,
		Specialties:         specialties,
		Certifications:      certifications,
	}
}

func (c *Cheff) Update(
	cheff_image, cheff_name, phone_number, address, locationId,
	locationPreference, cookingExperience, specialties, certifications,
	registration_status string,
) {
	c.Cheff_image = cheff_image
	c.Cheff_name = cheff_name
	c.Phone_number = phone_number
	c.Address = address
	c.LocationID = locationId
	c.Registration_status = registration_status
	c.LocationPreference = locationPreference
	c.CookingExperience = cookingExperience
	c.Specialties = specialties
	c.Certifications = certifications
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
		return errors.New("Chef not found")
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
