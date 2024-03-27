package domain

import "github.com/google/uuid"

type Order struct {
	ID               string      `json:"order_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID       string      `json:"customer_id"`
	ChefID           string      `json:"chef_id"`
	Items            []OrderItem `json:"items"`
	Status           string      `json:"status"`
	Order_Date       string      `json:"order_date"`
	DriverID         string      `json:"driver_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Delivery_address string      `json:"delivery_address"`
}

func NewOrder(customerID, chefID string, items []OrderItem, status, delivery_address string) (*Order, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Order{
		ID:               id.String(),
		CustomerID:       customerID,
		ChefID:           chefID,
		Items:            items,
		Status:           status,
		Delivery_address: delivery_address,
	}, nil
}
