package usecase

import "evaeats/user-service/internal/order/entity"

type CreateOrderItemInputDto struct {
	OrderID         string  `json:"order_id"`
	DishID          string  `json:"dish_id"`
	UnitPrice       float64 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	ChefID          string  `json:"chef_id"`
	DeliveryAddress string  `json:"delivery_address"`
	Note            string  `json:"note"`
}

type CreateOrderItemOutputDto struct {
	ID              string  `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID         string  `json:"order_id"`
	DishID          string  `json:"dish_id"`
	UnitPrice       float64 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	ChefID          string  `json:"chef_id"`
	DeliveryAddress string  `json:"delivery_address"`
	Note            string  `json:"note"`
}

type CreateOrderItemUseCase struct {
	OrderItemRepository entity.OrderItemRepository
}

func NewCreateOrderItemUseCase(orderItemRepository entity.OrderItemRepository) *CreateOrderItemUseCase {
	return &CreateOrderItemUseCase{OrderItemRepository: orderItemRepository}
}

func (u *CreateOrderItemUseCase) Execute(input CreateOrderItemInputDto) (*CreateOrderItemOutputDto, error) {
	newOrderItem := entity.NewOrderItem(
		input.OrderID,
		input.DishID,
		input.ChefID,
		input.UnitPrice,
		input.Quantity,
		input.DeliveryAddress,
		input.Note,
	)

	err := u.OrderItemRepository.Create(newOrderItem)
	if err != nil {
		return nil, err
	}

	output := &CreateOrderItemOutputDto{
		ID:              newOrderItem.ID,
		OrderID:         newOrderItem.OrderID,
		DishID:          newOrderItem.DishID,
		UnitPrice:       newOrderItem.UnitPrice,
		Quantity:        newOrderItem.Quantity,
		ChefID:          newOrderItem.ChefID,
		DeliveryAddress: newOrderItem.DeliveryAddress,
		Note:            newOrderItem.Note,
	}

	return output, nil
}

type UpdateOrderItemInputDto struct {
	ID              string  `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID         string  `json:"order_id"`
	DishID          string  `json:"dish_id"`
	UnitPrice       float64 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	ChefID          string  `json:"chef_id"`
	DeliveryAddress string  `json:"delivery_address"`
	Note            string  `json:"note"`
}

type UpdateOrderItemOutputDto struct {
	ID              string  `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID         string  `json:"order_id"`
	DishID          string  `json:"dish_id"`
	UnitPrice       float64 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	ChefID          string  `json:"chef_id"`
	DeliveryAddress string  `json:"delivery_address"`
	Note            string  `json:"note"`
}

type DeleteOrderItemInputDto struct {
	ID string `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
}

type GetOrderItemInputDto struct {
	ID string `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
}

type GetOrderItemOutputDto struct {
	ID              string  `json:"order_item_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID         string  `json:"order_id"`
	DishID          string  `json:"dish_id"`
	UnitPrice       float64 `json:"unit_price"`
	Quantity        int     `json:"quantity"`
	ChefID          string  `json:"chef_id"`
	DeliveryAddress string  `json:"delivery_address"`
	Note            string  `json:"note"`
}

type UpdateOrderItemUseCase struct {
	OrderItemRepository entity.OrderItemRepository
}

func NewUpdateOrderItemUseCase(orderItemRepository entity.OrderItemRepository) *UpdateOrderItemUseCase {
	return &UpdateOrderItemUseCase{OrderItemRepository: orderItemRepository}
}

func (u *UpdateOrderItemUseCase) Execute(input UpdateOrderItemInputDto) (*UpdateOrderItemOutputDto, error) {
	existingOrderItem, err := u.OrderItemRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existingOrderItem.OrderID = input.OrderID
	existingOrderItem.DishID = input.DishID
	existingOrderItem.UnitPrice = input.UnitPrice
	existingOrderItem.Quantity = input.Quantity
	existingOrderItem.ChefID = input.ChefID
	existingOrderItem.DeliveryAddress = input.DeliveryAddress
	existingOrderItem.Note = input.Note

	err = u.OrderItemRepository.Update(existingOrderItem)
	if err != nil {
		return nil, err
	}

	output := &UpdateOrderItemOutputDto{
		ID:              existingOrderItem.ID,
		OrderID:         existingOrderItem.OrderID,
		DishID:          existingOrderItem.DishID,
		UnitPrice:       existingOrderItem.UnitPrice,
		Quantity:        existingOrderItem.Quantity,
		ChefID:          existingOrderItem.ChefID,
		DeliveryAddress: existingOrderItem.DeliveryAddress,
		Note:            existingOrderItem.Note,
	}

	return output, nil
}

type DeleteOrderItemUseCase struct {
	OrderItemRepository entity.OrderItemRepository
}

func NewDeleteOrderItemUseCase(orderItemRepository entity.OrderItemRepository) *DeleteOrderItemUseCase {
	return &DeleteOrderItemUseCase{OrderItemRepository: orderItemRepository}
}

func (u *DeleteOrderItemUseCase) Execute(input DeleteOrderItemInputDto) error {
	return u.OrderItemRepository.DeleteByID(input.ID)
}

type GetOrderItemUseCase struct {
	OrderItemRepository entity.OrderItemRepository
}

func NewGetOrderItemUseCase(orderItemRepository entity.OrderItemRepository) *GetOrderItemUseCase {
	return &GetOrderItemUseCase{OrderItemRepository: orderItemRepository}
}

func (u *GetOrderItemUseCase) Execute(input GetOrderItemInputDto) (*GetOrderItemOutputDto, error) {
	orderItem, err := u.OrderItemRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	output := &GetOrderItemOutputDto{
		ID:              orderItem.ID,
		OrderID:         orderItem.OrderID,
		DishID:          orderItem.DishID,
		UnitPrice:       orderItem.UnitPrice,
		Quantity:        orderItem.Quantity,
		ChefID:          orderItem.ChefID,
		DeliveryAddress: orderItem.DeliveryAddress,
		Note:            orderItem.Note,
	}

	return output, nil
}
