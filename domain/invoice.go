package domain

type Invoice struct {
	ID             string  `json:"invoice_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OrderID        string  `json:"order_id"`
	Payment_method string  `json:"payment_method" validate:"eq=CARD|eq=MOBILEMONEY"`
	Payment_status string  `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_date   string  `json:"payment_date"`
	UserID         string  `json:"customer_id"`
	Amount         float64 `json:"amount"`
}
