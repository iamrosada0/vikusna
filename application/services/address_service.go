package services

import (
	"evaeats/application/repositories"
	"evaeats/domain"
)

type AddressService struct {
	AddressRepository repositories.AddressRepository
}

func NewAddressService(addressRepo repositories.AddressRepository) *AddressService {
	return &AddressService{
		AddressRepository: addressRepo,
	}
}

func (s *AddressService) CreateAddress(customerID, street, city, state, postalCode string) (*domain.Address, error) {
	// Validate input data, if necessary

	// Create the address in the database
	newAddress, err := s.AddressRepository.Insert(customerID, street, city, state, postalCode)
	if err != nil {
		return nil, err
	}

	return newAddress, nil
}

func (s *AddressService) GetAddressByID(id string) (*domain.Address, error) {
	// Fetch the address by ID from the database
	address, err := s.AddressRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return address, nil
}
