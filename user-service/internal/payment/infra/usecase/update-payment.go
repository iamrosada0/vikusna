package usecase

import "evaeats/user-service/internal/payment/entity"

type UpdatePaymentInputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdatePaymentOutputDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdatePaymentUseCase struct {
	PaymentRepository entity.PaymentRepository
}

func NewUpdatePaymentUseCase(PaymentRepository entity.PaymentRepository) *UpdatePaymentUseCase {
	return &UpdatePaymentUseCase{PaymentRepository: PaymentRepository}
}

func (u *UpdatePaymentUseCase) Execute(input UpdatePaymentInputDto) (*UpdatePaymentOutputDto, error) {
	existingPayment, err := u.PaymentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	existingPayment.Update(
		input.Name,
		input.Email,
	)

	err = u.PaymentRepository.Update(existingPayment)
	if err != nil {
		return nil, err
	}

	return &UpdatePaymentOutputDto{
		ID:    existingPayment.ID,
		Name:  existingPayment.Name,
		Email: existingPayment.Email,
	}, nil
}
