package domain

import "github.com/google/uuid"

// Transação de Pagamento representa uma transação de pagamento entre um cliente e um chef
type PaymentTransaction struct {
	ID            string  `json:"payment_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

func NewPaymentTransaction(orderID, paymentMethod string, amount float64) (*PaymentTransaction, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &PaymentTransaction{
		ID:            id.String(),
		OrderID:       orderID,
		Amount:        amount,
		PaymentMethod: paymentMethod,
	}, nil
}
