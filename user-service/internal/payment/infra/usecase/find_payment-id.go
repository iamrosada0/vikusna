package usecase

import "evaeats/user-service/internal/payment/entity"

type GetPaymentByIDInputDto struct {
	ID uint `json:"id"`
}

type GetPaymentByIDOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetPaymentByIDUseCase struct {
	PaymentRepository entity.PaymentRepository
}

func NewGetPaymentByIDUseCase(PaymentRepository entity.PaymentRepository) *GetPaymentByIDUseCase {
	return &GetPaymentByIDUseCase{PaymentRepository: PaymentRepository}
}

func (u *GetPaymentByIDUseCase) Execute(input GetPaymentByIDInputDto) (*GetPaymentByIDOutputDto, error) {
	Payment, err := u.PaymentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetPaymentByIDOutputDto{
		ID:    Payment.ID,
		Name:  Payment.Name,
		Email: Payment.Email,
	}, nil
}
