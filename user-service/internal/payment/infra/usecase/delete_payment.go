package usecase

import "evaeats/user-service/internal/payment/entity"

type DeletePaymentInputDto struct {
	ID uint `json:"id"`
}

type DeletePaymentOutputDto struct {
	ID uint `json:"id"`
}

type DeletePaymentUseCase struct {
	PaymentRepository entity.PaymentRepository
}

func NewDeletePaymentUseCase(PaymentRepository entity.PaymentRepository) *DeletePaymentUseCase {
	return &DeletePaymentUseCase{PaymentRepository: PaymentRepository}
}

func (u *DeletePaymentUseCase) Execute(input DeletePaymentInputDto) (*DeletePaymentOutputDto, error) {
	err := u.PaymentRepository.DeleteByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &DeletePaymentOutputDto{ID: input.ID}, nil
}
