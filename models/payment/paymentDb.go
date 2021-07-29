package payment

import "time"

type Payment struct {
	PaymentId    int       `json:"paymentId" gorm:"primaryKey"`
	TotalPayment int       `json:"totalPayment"`
	IsPaid       bool      `json:"isPaid"`
	CartId       int       `json:"cartId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
