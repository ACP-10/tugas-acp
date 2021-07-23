package cart

import "time"

type Cart struct {
	CartId     int `json:"cartId"`
	IsCheckout bool `json:"isCheckout"`
	CustomerId int `json:"customerId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

}