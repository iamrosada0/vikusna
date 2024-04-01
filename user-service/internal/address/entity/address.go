package entity

import (
	"errors"

	"github.com/google/uuid"
)

// AddressRepository defines the methods for interacting with address data
type AddressRepository interface {
	Create(address *Address) error
	FindAll() ([]*Address, error)
	Update(address *Address) error
	DeleteByID(id string) error
	GetByID(id string) (*Address, error)
}

// Address represents the address entity
type Address struct {
	ID         string `json:"address_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID string `json:"customer_id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}

// NewAddress creates a new address instance
func NewAddress(customerID, street, city, state, postalCode string) *Address {
	return &Address{
		ID:         uuid.New().String(),
		CustomerID: customerID,
		Street:     street,
		City:       city,
		State:      state,
		PostalCode: postalCode,
	}
}

// InMemoryAddressRepository is an in-memory repository for addresses
type InMemoryAddressRepository struct {
	Addresses map[string]*Address
}

// NewInMemoryAddressRepository creates a new instance of InMemoryAddressRepository
func NewInMemoryAddressRepository() *InMemoryAddressRepository {
	return &InMemoryAddressRepository{
		Addresses: make(map[string]*Address),
	}
}

// Create creates a new address in the repository
func (r *InMemoryAddressRepository) Create(address *Address) error {
	if _, exists := r.Addresses[address.ID]; exists {
		return errors.New("address already exists")
	}
	r.Addresses[address.ID] = address
	return nil
}

// DeleteByID deletes an address by its ID from the repository
func (r *InMemoryAddressRepository) DeleteByID(id string) error {
	if _, exists := r.Addresses[id]; !exists {
		return errors.New("address not found")
	}
	delete(r.Addresses, id)
	return nil
}

// FindAll retrieves all addresses from the repository
func (r *InMemoryAddressRepository) FindAll() ([]*Address, error) {
	var allAddresses []*Address
	for _, address := range r.Addresses {
		allAddresses = append(allAddresses, address)
	}
	return allAddresses, nil
}

// Update updates an address in the repository
func (r *InMemoryAddressRepository) Update(address *Address) error {
	if _, exists := r.Addresses[address.ID]; !exists {
		return errors.New("address not found")
	}
	r.Addresses[address.ID] = address
	return nil
}

// GetByID retrieves an address by its ID from the repository
func (r *InMemoryAddressRepository) GetByID(id string) (*Address, error) {
	if address, exists := r.Addresses[id]; exists {
		return address, nil
	}
	return nil, errors.New("address not found")
}
