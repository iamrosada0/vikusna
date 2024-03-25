package domain

import "github.com/google/uuid"

// Endereço representa o endereço de entrega associado a um pedido
type Address struct {
	ID         string `json:"id_address" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID string
	Street     string
	City       string
	State      string
	PostalCode string
}

func NewAddress(customerID, street, city, state, postalCode string) (*Address, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Address{
		ID:         id.String(),
		CustomerID: customerID,
		Street:     street,
		City:       city,
		State:      state,
		PostalCode: postalCode,
	}, nil
}
