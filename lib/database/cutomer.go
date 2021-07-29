package database

import (
	"tugas-acp/configs"
	"tugas-acp/models/customer"
)

func GetCustomer(email string) (dataresult []customer.Customer, err error) {
	err = configs.DB.First(&dataresult, "email", email).Error
	if err != nil {
		return nil, err
	}
	return
}
