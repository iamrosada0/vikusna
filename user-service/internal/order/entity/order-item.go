package entity

import (
	"errors"

	"github.com/google/uuid"
)

type OrderItemRepository interface {
	Create(orderItem *OrderItem) error
	//FindAll() ([]*OrderItem, error)
	Update(orderItem *OrderItem) error
	DeleteByID(id string) error
	GetByID(id string) (*OrderItem, error)
}

type OrderItem struct {
	ID              string  `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID         string  `json:"order_id"`
	DishID          string  `json:"dish_id"`
	UnitPrice       float64 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	ChefID          string  `json:"chef_id"`
	DeliveryAddress string  `json:"delivery_address"`
	Note            string  `json:"note"`
}

func NewOrderItem(orderID, dishID, chefID string, unitPrice float64, quantity int, deliveryAddress, note string) *OrderItem {
	return &OrderItem{
		ID:              uuid.New().String(),
		OrderID:         orderID,
		DishID:          dishID,
		UnitPrice:       unitPrice,
		Quantity:        quantity,
		ChefID:          chefID,
		DeliveryAddress: deliveryAddress,
		Note:            note,
	}
}

type InMemoryOrderItemRepository struct {
	OrderItems map[string]*OrderItem
}

func NewInMemoryOrderItemRepository() *InMemoryOrderItemRepository {
	return &InMemoryOrderItemRepository{
		OrderItems: make(map[string]*OrderItem),
	}
}

func (r *InMemoryOrderItemRepository) Create(orderItem *OrderItem) error {
	if _, exists := r.OrderItems[orderItem.ID]; exists {
		return errors.New("order item already exists")
	}
	r.OrderItems[orderItem.ID] = orderItem
	return nil
}

func (r *InMemoryOrderItemRepository) DeleteByID(id string) error {
	if _, exists := r.OrderItems[id]; !exists {
		return errors.New("order item not found")
	}
	delete(r.OrderItems, id)
	return nil
}

// func (r *InMemoryOrderItemRepository) FindAll() ([]*OrderItem, error) {
// 	var allOrderItems []*OrderItem
// 	for _, orderItem := range r.OrderItems {
// 		allOrderItems = append(allOrderItems, orderItem)
// 	}
// 	return allOrderItems, nil
// }

func (r *InMemoryOrderItemRepository) Update(orderItem *OrderItem) error {
	if _, exists := r.OrderItems[orderItem.ID]; !exists {
		return errors.New("order item not found")
	}
	r.OrderItems[orderItem.ID] = orderItem
	return nil
}

func (r *InMemoryOrderItemRepository) GetByID(id string) (*OrderItem, error) {
	if orderItem, exists := r.OrderItems[id]; exists {
		return orderItem, nil
	}
	return nil, errors.New("order item not found")
}
