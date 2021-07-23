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