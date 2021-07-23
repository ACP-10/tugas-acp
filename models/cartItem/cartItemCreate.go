package cartitem

type CartItemCreate struct {
	CartId    int       `json:"cartId"`
	Quantity  int       `json:"quantity"`
}