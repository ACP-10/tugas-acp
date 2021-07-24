package database

import (
	"tugas-acp/configs"
	"tugas-acp/models/cart")


func GetCartAll() (dataresult []cart.Cart, err error) {
	err = configs.DB.Find(&dataresult).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetCartByCustomer(customerId int, isCheckout bool)(dataresult []cart.CartResult, err error) {
	err = configs.DB.Model(&cart.Cart{}).Select("carts.cart_id, carts.is_checkout, customers.name AS customer").
	Joins("JOIN customers ON customers.customer_id = carts.customer_id").
	Where("carts.customer_id = ? AND carts.is_checkout = ?",customerId, isCheckout).
	Scan(&dataresult).Error
	if err != nil {
		return nil, err
	}
	return
}

