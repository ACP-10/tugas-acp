package cart

type CartCreate struct {
	CustomerId int `json:"customerId"`
	IsCheckout bool `json:"isCheckout"`
}