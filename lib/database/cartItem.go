package database

import (
	"tugas-acp/configs"
	cartitem "tugas-acp/models/cartItem"
)


func GetCartItemByCartId(cartId int) (dataresult []cartitem.CartItemResult, err error) {
	err = configs.DB.Model(&cartitem.CartItem{}).
	Select("products.product_name AS product, categories.category_name AS category,cart_items.item_id, cart_items.cart_id,cart_items.quantity,cart_items.created_at,cart_items.updated_at").
	Joins("JOIN products ON products.product_id = cart_items.product_id").
	Joins("JOIN categories ON categories.category_id = products.category_id").
	Where("cart_id = ?", cartId).Scan(&dataresult).Error

	if err != nil {
		return nil, err
	}
	return
}