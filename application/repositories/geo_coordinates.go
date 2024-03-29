package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type GeoCoordinatesRepository interface {
	Insert(name string, latitude, longitude float64) (*domain.GeoCoordinates, error)
	Find(id string) (*domain.GeoCoordinates, error)
}

type GeoCoordinatesRepositoryDb struct {
	Db *gorm.DB
}

func (repo GeoCoordinatesRepositoryDb) Insert(name string, latitude, longitude float64) (*domain.GeoCoordinates, error) {
	// Generate a new GeoCoordinates with a generated ID
	newGeoCoordinates, err := domain.NewGeoCoordinates(name, latitude, longitude)
	if err != nil {
		return nil, err
	}

	// Insert the GeoCoordinates into the database
	if err := repo.Db.Create(newGeoCoordinates).Error; err != nil {
		return nil, err
	}

	return newGeoCoordinates, nil
}

func (repo GeoCoordinatesRepositoryDb) Find(id string) (*domain.GeoCoordinates, error) {
	var geoCoordinates domain.GeoCoordinates
	if err := repo.Db.First(&geoCoordinates, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("geo coordinates with ID %s not found", id)
		}
		return nil, err
	}
	return &geoCoordinates, nil
}
