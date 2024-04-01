package usecase

import "evaeats/user-service/internal/order/entity"

type GetOrderByIDInputDto struct {
	ID string `json:"id"`
}

type GetOrderByIDOutputDto struct {
	ID string `json:"order_id" valid:"uuid" gorm:"type:uuid;primary_key"`

	CustomerID      string             `json:"customer_id"`
	ChefID          string             `json:"chef_id"`
	Items           []entity.OrderItem `json:"items"`
	Status          string             `json:"status"`
	OrderDate       string             `json:"order_date"`
	DriverID        string             `json:"driver_id" valid:"uuid" gorm:"type:uuid"`
	DeliveryAddress string             `json:"delivery_address"`
}

type GetOrderByIDUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewGetOrderByIDUseCase(OrderRepository entity.OrderRepository) *GetOrderByIDUseCase {
	return &GetOrderByIDUseCase{OrderRepository: OrderRepository}
}

func (u *GetOrderByIDUseCase) Execute(input GetOrderByIDInputDto) (*GetOrderByIDOutputDto, error) {
	order, err := u.OrderRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetOrderByIDOutputDto{
		ID:              order.ID,
		CustomerID:      order.CustomerID,
		ChefID:          order.ChefID,
		Items:           order.Items,
		Status:          order.Status,
		OrderDate:       order.OrderDate,
		DriverID:        order.DriverID,
		DeliveryAddress: order.DeliveryAddress,
	}, nil
}
