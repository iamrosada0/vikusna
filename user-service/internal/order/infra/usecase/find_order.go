package usecase

import "evaeats/user-service/internal/order/entity"

type GetAllOrdersOutputDto struct {
	Orders []*entity.Order `json:"Orders"`
}

type GetAllOrdersUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewGetAllOrdersUseCase(OrderRepository entity.OrderRepository) *GetAllOrdersUseCase {
	return &GetAllOrdersUseCase{OrderRepository: OrderRepository}
}

func (u *GetAllOrdersUseCase) Execute() (*GetAllOrdersOutputDto, error) {
	Orders, err := u.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &GetAllOrdersOutputDto{Orders: Orders}, nil
}
