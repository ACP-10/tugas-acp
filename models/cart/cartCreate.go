package cart

type CartCreate struct {
	IsCheckout bool `json:"isCheckout"`
	CustomerId int `json:"customerId"`
}