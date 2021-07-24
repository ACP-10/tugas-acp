package cartitem

import "time"

type CartItem struct {
	ItemId    int       `json:"itemId" gorm:"primaryKey"`
	CartId    int       `json:"cartId"`
	ProductId int 		`json:"productId"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}