package cartitem

import "time"

type CartItem struct {
	ItemId    int       `json:"itemId"`
	CartId    int       `json:"cartId"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}