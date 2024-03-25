package domain

import "github.com/google/uuid"

type Order struct {
	ID         string `json:"order_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID string
	ChefID     string
	Items      []OrderItem
	Status     string
	Address    string
}

func NewOrder(customerID, chefID string, items []OrderItem, status, address string) (*Order, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Order{
		ID:         id.String(),
		CustomerID: customerID,
		ChefID:     chefID,
		Items:      items,
		Status:     status,
		Address:    address,
	}, nil
}
