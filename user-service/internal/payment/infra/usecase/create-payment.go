package usecase

import "evaeats/user-service/internal/payment/entity"

type CreatePaymentInputDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreatePaymentOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreatePaymentUseCase struct {
	PaymentRepository entity.PaymentRepository
}

func NewCreatePaymentUseCase(PaymentRepository entity.PaymentRepository) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{PaymentRepository: PaymentRepository}
}

func (u *CreatePaymentUseCase) Execute(input CreatePaymentInputDto) (*CreatePaymentOutputDto, error) {
	Payment := entity.NewPayment(
		input.Name,
		input.Email,
	)

	err := u.PaymentRepository.Create(Payment)
	if err != nil {
		return nil, err
	}

	return &CreatePaymentOutputDto{
		ID:    Payment.ID,
		Name:  Payment.Name,
		Email: Payment.Email,
	}, nil
}
