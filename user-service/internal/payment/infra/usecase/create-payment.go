package usecase

import "evaeats/user-service/internal/payment/entity"

type CreatePaymentInputDto struct {
	OrderID       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type CreatePaymentOutputDto struct {
	ID            string  `json:"payment_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type CreatePaymentUseCase struct {
	PaymentRepository entity.PaymentRepository
}

func NewCreatePaymentUseCase(paymentRepository entity.PaymentRepository) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{PaymentRepository: paymentRepository}
}

func (u *CreatePaymentUseCase) Execute(input CreatePaymentInputDto) (*CreatePaymentOutputDto, error) {
	newPayment, err := entity.NewPaymentTransaction(
		input.OrderID,
		input.PaymentMethod,
		input.Amount,
	)
	if err != nil {
		return nil, err
	}

	err = u.PaymentRepository.Create(newPayment)
	if err != nil {
		return nil, err
	}

	output := &CreatePaymentOutputDto{
		ID:            newPayment.ID,
		OrderID:       newPayment.OrderID,
		Amount:        newPayment.Amount,
		PaymentMethod: newPayment.PaymentMethod,
	}

	return output, nil
}
