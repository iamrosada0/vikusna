package domain

import (
	"github.com/google/uuid"
)

// Item do Pedido representa um item específico dentro de um pedido
type OrderItem struct {
	ID       string `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID  string `json:"order_id"`
	DishID   string `json:"dish_id"`
	Quantity int    `json:"quantity"`
}

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
