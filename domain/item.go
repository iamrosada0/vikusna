package domain

import (
	"github.com/google/uuid"
)

// OrderItem represents a specific item within an order
type OrderItem struct {
	ID              string  `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID         string  `json:"order_id"`
	DishID          string  `json:"dish_id"`
	UnitPrice       float64 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	HomeChefEats_id string  `json:"homeChefEats_id"` // Represents the restaurant (homeChefEats_id)
}

// NewOrderItem creates a new OrderItem instance
func NewOrderItem(orderID, dishID string, quantity int) (*OrderItem, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &OrderItem{
		ID:       id.String(),
		OrderID:  orderID,
		DishID:   dishID,
		Quantity: quantity,
	}, nil
}
