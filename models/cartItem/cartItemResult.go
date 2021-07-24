package cartitem

type CartItemResult struct{
	CartId int `json:"cartId"`
	ItemId int `json:"itemId"`
	Product string `json:"product"`
	Category string `json:"category"`
	Quantity int `json:"quantity"`
}