package payment

import "time"

type Payment struct {
	PaymentId    int       `json:"paymentId"`
	TotalPayment int       `json:"totalPayment"`
	CartId       int       `json:"cartId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}