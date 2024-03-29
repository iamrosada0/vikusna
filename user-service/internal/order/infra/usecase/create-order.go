package usecase

import "evaeats/user-service/internal/order/entity"

type CreateOrderInputDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateOrderOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewCreateOrderUseCase(OrderRepository entity.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{OrderRepository: OrderRepository}
}

func (u *CreateOrderUseCase) Execute(input CreateOrderInputDto) (*CreateOrderOutputDto, error) {
	Order := entity.NewOrder(
		input.Name,
		input.Email,
	)

	err := u.OrderRepository.Create(Order)
	if err != nil {
		return nil, err
	}

	return &CreateOrderOutputDto{
		ID:    Order.ID,
		Name:  Order.Name,
		Email: Order.Email,
	}, nil
}
