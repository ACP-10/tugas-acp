package cart

import cartitem "tugas-acp/models/cartItem"

type CartResult struct {
	CartId       int  `json:"cartId"`
	Customer string  `json:"customer"`
	IsCheckout   bool `json:"isCheckout"`
	Item         []cartitem.CartItemResult `json:"item" gorm:"foreignKey:CartId;references:CartId"`
}