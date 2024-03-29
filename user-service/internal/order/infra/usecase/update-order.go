package usecase

import "evaeats/user-service/internal/order/entity"

type UpdateOrderInputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateOrderOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateOrderUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewUpdateOrderUseCase(OrderRepository entity.OrderRepository) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{OrderRepository: OrderRepository}
}

func (u *UpdateOrderUseCase) Execute(input UpdateOrderInputDto) (*UpdateOrderOutputDto, error) {
	existingOrder, err := u.OrderRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existingOrder.Update(
		input.Name,
		input.Email,
	)

	err = u.OrderRepository.Update(existingOrder)
	if err != nil {
		return nil, err
	}

	return &UpdateOrderOutputDto{
		ID:    existingOrder.ID,
		Name:  existingOrder.Name,
		Email: existingOrder.Email,
	}, nil
}
