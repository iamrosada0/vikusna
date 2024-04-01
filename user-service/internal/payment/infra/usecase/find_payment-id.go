package usecase

import "evaeats/user-service/internal/payment/entity"

type GetPaymentByIDInputDto struct {
	ID string `json:"payment_id" valid:"uuid" gorm:"type:uuid;primary_key"`
}

type GetPaymentByIDOutputDto struct {
	ID            string  `json:"payment_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type GetPaymentByIDUseCase struct {
	PaymentRepository entity.PaymentRepository
}

func NewGetPaymentByIDUseCase(paymentRepository entity.PaymentRepository) *GetPaymentByIDUseCase {
	return &GetPaymentByIDUseCase{PaymentRepository: paymentRepository}
}

func (u *GetPaymentByIDUseCase) Execute(input GetPaymentByIDInputDto) (*GetPaymentByIDOutputDto, error) {
	payment, err := u.PaymentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetPaymentByIDOutputDto{
		ID:            payment.ID,
		OrderID:       payment.OrderID,
		Amount:        payment.Amount,
		PaymentMethod: payment.PaymentMethod,
	}, nil
}
