package usecase

import "evaeats/user-service/internal/address/entity"

type AddressUseCase struct {
	addressRepository entity.AddressRepository
}

func NewAddressUseCase(addressRepository entity.AddressRepository) *AddressUseCase {
	return &AddressUseCase{addressRepository: addressRepository}
}

func (uc *AddressUseCase) CreateAddress(customerID, street, city, state, postalCode string) (*entity.Address, error) {
	address := entity.NewAddress(customerID, street, city, state, postalCode)

	if err := uc.addressRepository.Create(address); err != nil {
		return nil, err
	}
	return address, nil
}

func (uc *AddressUseCase) UpdateAddress(id, customerID, street, city, state, postalCode string) error {
	address := &entity.Address{
		ID:         id,
		CustomerID: customerID,
		Street:     street,
		City:       city,
		State:      state,
		PostalCode: postalCode,
	}
	return uc.addressRepository.Update(address)
}

func (uc *AddressUseCase) DeleteAddress(id string) error {
	return uc.addressRepository.DeleteByID(id)
}

func (uc *AddressUseCase) GetAddressByID(id string) (*entity.Address, error) {
	return uc.addressRepository.GetByID(id)
}
