package domain

import "github.com/google/uuid"

type Order struct {
	ID         string      `json:"order_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID string      `json:"customer_id"`
	ChefID     string      `json:"chef_id"`
	Items      []OrderItem `json:"items"`
	Status     string      `json:"status"`
	Address    string      `json:"address"`
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
