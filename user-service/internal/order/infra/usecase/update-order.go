package usecase

import "evaeats/user-service/internal/order/entity"

type UpdateOrderInputDto struct {
	ID              string             `json:"order_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID      string             `json:"customer_id"`
	ChefID          string             `json:"chef_id"`
	Items           []entity.OrderItem `json:"items"`
	Status          string             `json:"status"`
	OrderDate       string             `json:"order_date"`
	DriverID        string             `json:"driver_id" valid:"uuid" gorm:"type:uuid"`
	DeliveryAddress string             `json:"delivery_address"`
}

type UpdateOrderOutputDto struct {
	ID              string             `json:"order_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	CustomerID      string             `json:"customer_id"`
	ChefID          string             `json:"chef_id"`
	Items           []entity.OrderItem `json:"items"`
	Status          string             `json:"status"`
	OrderDate       string             `json:"order_date"`
	DriverID        string             `json:"driver_id" valid:"uuid" gorm:"type:uuid"`
	DeliveryAddress string             `json:"delivery_address"`
}

type UpdateOrderUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewUpdateOrderUseCase(orderRepository entity.OrderRepository) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{OrderRepository: orderRepository}
}

func (u *UpdateOrderUseCase) Execute(input UpdateOrderInputDto) (*UpdateOrderOutputDto, error) {
	existingOrder, err := u.OrderRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Update the fields of the existing order with the values from the input DTO
	existingOrder.CustomerID = input.CustomerID
	existingOrder.ChefID = input.ChefID
	existingOrder.Items = input.Items
	existingOrder.Status = input.Status
	existingOrder.OrderDate = input.OrderDate
	existingOrder.DriverID = input.DriverID
	existingOrder.DeliveryAddress = input.DeliveryAddress

	err = u.OrderRepository.Update(existingOrder)
	if err != nil {
		return nil, err
	}

	return &UpdateOrderOutputDto{
		ID:              existingOrder.ID,
		CustomerID:      existingOrder.CustomerID,
		ChefID:          existingOrder.ChefID,
		Items:           existingOrder.Items,
		Status:          existingOrder.Status,
		OrderDate:       existingOrder.OrderDate,
		DriverID:        existingOrder.DriverID,
		DeliveryAddress: existingOrder.DeliveryAddress,
	}, nil
}
