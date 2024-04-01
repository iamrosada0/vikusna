package services

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type GeoCoordinatesService struct {
	GeoCoordinatesRepository repositories.GeoCoordinatesRepository
}

func NewGeoCoordinatesService(geoCoordinatesRepo repositories.GeoCoordinatesRepository) *GeoCoordinatesService {
	return &GeoCoordinatesService{
		GeoCoordinatesRepository: geoCoordinatesRepo,
	}
}

func (s *GeoCoordinatesService) CreateGeoCoordinates(name string, latitude, longitude float64) (*domain.GeoCoordinates, error) {
	// Validate input data, if necessary

	// Create the geo coordinates in the database
	newGeoCoordinates, err := s.GeoCoordinatesRepository.Insert(name, latitude, longitude)
	if err != nil {
		return nil, err
	}

	return newGeoCoordinates, nil
}

func (s *GeoCoordinatesService) GetGeoCoordinatesByID(id string) (*domain.GeoCoordinates, error) {
	// Fetch the geo coordinates by ID from the database
	geoCoordinates, err := s.GeoCoordinatesRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return geoCoordinates, nil
}
