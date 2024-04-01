package usecase

import "evaeats/user-service/internal/order/entity"

type CreateOrderInputDto struct {
	CustomerID      string             `json:"customer_id"`
	ChefID          string             `json:"chef_id"`
	Items           []entity.OrderItem `json:"items"`
	Status          string             `json:"status"`
	OrderDate       string             `json:"order_date"`
	DriverID        string             `json:"driver_id" valid:"uuid" gorm:"type:uuid"`
	DeliveryAddress string             `json:"delivery_address"`
}

type CreateOrderOutputDto struct {
	ID              string             `json:"order_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID      string             `json:"customer_id"`
	ChefID          string             `json:"chef_id"`
	Items           []entity.OrderItem `json:"items"`
	Status          string             `json:"status"`
	OrderDate       string             `json:"order_date"`
	DriverID        string             `json:"driver_id" valid:"uuid" gorm:"type:uuid"`
	DeliveryAddress string             `json:"delivery_address"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{OrderRepository: orderRepository}
}

func (u *CreateOrderUseCase) Execute(input CreateOrderInputDto) (*CreateOrderOutputDto, error) {
	newOrder := entity.NewOrder(
		input.CustomerID,
		input.ChefID,
		input.Status,
		input.DriverID,
		input.DeliveryAddress,
		input.Items,
	)

	err := u.OrderRepository.Create(newOrder)
	if err != nil {
		return nil, err
	}

	output := &CreateOrderOutputDto{
		ID:              newOrder.ID,
		CustomerID:      newOrder.CustomerID,
		ChefID:          newOrder.ChefID,
		Items:           newOrder.Items,
		Status:          newOrder.Status,
		OrderDate:       newOrder.OrderDate,
		DriverID:        newOrder.DriverID,
		DeliveryAddress: newOrder.DeliveryAddress,
	}

	return output, nil
}
