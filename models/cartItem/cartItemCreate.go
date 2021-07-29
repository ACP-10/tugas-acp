package cartitem

type CartItemCreate struct {
	CartId    int       `json:"cartId"`
	ProductId int 		`json:"productId"`
	Quantity  int       `json:"quantity"`
}