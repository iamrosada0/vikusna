package repository

import (
	"errors"
	"evaeats/user-service/internal/address/entity"
	"sync"
)

type AddressRepositoryInMemory struct {
	addresses map[string]*entity.Address
	mutex     sync.RWMutex
}

func NewAddressRepositoryInMemory() *AddressRepositoryInMemory {
	return &AddressRepositoryInMemory{
		addresses: make(map[string]*entity.Address),
	}
}

func (r *AddressRepositoryInMemory) Create(address *entity.Address) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.addresses[address.ID]; exists {
		return errors.New("address already exists")
	}
	r.addresses[address.ID] = address
	return nil
}

func (r *AddressRepositoryInMemory) Update(address *entity.Address) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.addresses[address.ID]; !exists {
		return errors.New("address not found")
	}
	r.addresses[address.ID] = address
	return nil
}

func (r *AddressRepositoryInMemory) DeleteByID(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.addresses[id]; !exists {
		return errors.New("address not found")
	}
	delete(r.addresses, id)
	return nil
}

func (r *AddressRepositoryInMemory) GetByID(id string) (*entity.Address, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	address, exists := r.addresses[id]
	if !exists {
		return nil, errors.New("address not found")
	}
	return address, nil
}
