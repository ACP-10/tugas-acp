package payment

type PaymentCreate struct {
	TotalPayment int       `json:"totalPayment"`
	CartId       int       `json:"cartId"`
}