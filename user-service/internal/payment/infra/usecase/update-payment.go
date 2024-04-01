package usecase

import "evaeats/user-service/internal/payment/entity"

type UpdatePaymentInputDto struct {
	ID            string  `json:"payment_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type UpdatePaymentOutputDto struct {
	ID            string  `json:"payment_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type UpdatePaymentUseCase struct {
	PaymentRepository entity.PaymentRepository
}

func NewUpdatePaymentUseCase(paymentRepository entity.PaymentRepository) *UpdatePaymentUseCase {
	return &UpdatePaymentUseCase{PaymentRepository: paymentRepository}
}

func (u *UpdatePaymentUseCase) Execute(input UpdatePaymentInputDto) (*UpdatePaymentOutputDto, error) {
	existingPayment, err := u.PaymentRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Update the existing payment with the new data
	existingPayment.OrderID = input.OrderID
	existingPayment.Amount = input.Amount
	existingPayment.PaymentMethod = input.PaymentMethod

	// Save the updated payment
	err = u.PaymentRepository.Update(existingPayment)
	if err != nil {
		return nil, err
	}

	// Return the updated payment details
	return &UpdatePaymentOutputDto{
		ID:            existingPayment.ID,
		OrderID:       existingPayment.OrderID,
		Amount:        existingPayment.Amount,
		PaymentMethod: existingPayment.PaymentMethod,
	}, nil
}
