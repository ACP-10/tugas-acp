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

func GetCartByCustomer(customerId int, isCheckout bool)(dataresult []cart.Cart, err error) {
	err = configs.DB.Where("customer_id = ? AND is_checkout = ?",customerId, isCheckout).First(&dataresult).Error
	if err != nil {
		return nil, err
	}
	return
}