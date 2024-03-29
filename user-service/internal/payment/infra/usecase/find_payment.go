package usecase

import "evaeats/user-service/internal/payment/entity"

type GetAllPaymentsOutputDto struct {
	Payments []*entity.Payment `json:"Payments"`
}

type GetAllPaymentsUseCase struct {
	PaymentRepository entity.PaymentRepository
}

func NewGetAllPaymentsUseCase(PaymentRepository entity.PaymentRepository) *GetAllPaymentsUseCase {
	return &GetAllPaymentsUseCase{PaymentRepository: PaymentRepository}
}

func (u *GetAllPaymentsUseCase) Execute() (*GetAllPaymentsOutputDto, error) {
	Payments, err := u.PaymentRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &GetAllPaymentsOutputDto{Payments: Payments}, nil
}
