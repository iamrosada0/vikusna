package usecase

import "evaeats/user-service/internal/order/entity"

type DeleteOrderInputDto struct {
	ID uint `json:"id"`
}

type DeleteOrderOutputDto struct {
	ID uint `json:"id"`
}

type DeleteOrderUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewDeleteOrderUseCase(OrderRepository entity.OrderRepository) *DeleteOrderUseCase {
	return &DeleteOrderUseCase{OrderRepository: OrderRepository}
}

func (u *DeleteOrderUseCase) Execute(input DeleteOrderInputDto) (*DeleteOrderOutputDto, error) {
	err := u.OrderRepository.DeleteByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteOrderOutputDto{ID: input.ID}, nil
}
