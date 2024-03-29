package usecase

import "evaeats/user-service/internal/order/entity"

type GetOrderByIDInputDto struct {
	ID uint `json:"id"`
}

type GetOrderByIDOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetOrderByIDUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewGetOrderByIDUseCase(OrderRepository entity.OrderRepository) *GetOrderByIDUseCase {
	return &GetOrderByIDUseCase{OrderRepository: OrderRepository}
}

func (u *GetOrderByIDUseCase) Execute(input GetOrderByIDInputDto) (*GetOrderByIDOutputDto, error) {
	Order, err := u.OrderRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetOrderByIDOutputDto{
		ID:    Order.ID,
		Name:  Order.Name,
		Email: Order.Email,
	}, nil
}
