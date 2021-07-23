package cart

type CartCreate struct {
	CartId     int `json:"cartId"`
	IsCheckout bool `json:"isCheckout"`
	CustomerId int `json:"customerId"`
}