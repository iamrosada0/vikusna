package repositories

import (
	"evaeats/domain"
	"fmt"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Insert(order_id, dish_id string, quantity int) (*domain.OrderItem, error)
	Find(id string) (*domain.OrderItem, error)
}

type OrderItemRepositoryDb struct {
	Db *gorm.DB
}

func (repo OrderItemRepositoryDb) Insert(order_id, dish_id string, quantity int) (*domain.OrderItem, error) {
	// Generate a new OrderItem with a generated ID
	newOrderItem, err := domain.NewOrderItem(order_id, dish_id, quantity)
	if err != nil {
		return nil, err
	}

	// Insert the OrderItem into the database
	if err := repo.Db.Create(newOrderItem).Error; err != nil {
		return nil, err
	}

	return newOrderItem, nil
}

func (repo OrderItemRepositoryDb) Find(id string) (*domain.OrderItem, error) {
	var OrderItem domain.OrderItem
	if err := repo.Db.First(&OrderItem, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("dish category with ID %s not found", id)
		}
		return nil, err
	}
	return &OrderItem, nil
}

func (repo *OrderItemRepositoryDb) Get() ([]map[string]interface{}, error) {
	var orderItems []*domain.OrderItem
	var orderItemHolder []map[string]interface{}

	// Fetch OrderItems along with associated Dish, Order, and Invoice
	if err := repo.Db.Preload("Dish").Preload("Order.Invoice").Find(&orderItems).Error; err != nil {
		return nil, err
	}

	for _, orderItem := range orderItems {
		var dish domain.Dish
		var order domain.Order
		var invoice domain.Invoice
		dishData := map[string]interface{}{
			"id":          dish.ID,
			"name":        dish.Name,
			"price":       dish.Price,
			"description": dish.Description,
			"available":   dish.Available,
		}

		orderData := map[string]interface{}{
			"id":               order.ID,
			"chef_id":          order.ChefID,
			"order_date":       order.Order_Date,
			"delivery_address": order.Delivery_address,
			"status":           order.Status,
			"customer_id":      order.CustomerID,
			"items":            order.Items,
		}

		invoiceData := map[string]interface{}{
			"id":             invoice.ID,
			"payment_date":   invoice.Payment_date,
			"payment_status": invoice.Payment_status,
		}

		orderItemData := map[string]interface{}{
			"id":              orderItem.ID,
			"quantity":        orderItem.Quantity,
			"unit_price":      orderItem.UnitPrice,
			"dish_details":    dishData,
			"homeChefEats_id": orderItem.HomeChefEats_id,
			"order_details":   orderData,
			"total_price":     float64(orderItem.Quantity) * orderItem.UnitPrice,
			"payment_details": invoiceData,
		}

		orderItemHolder = append(orderItemHolder, orderItemData)
	}

	return orderItemHolder, nil
}
