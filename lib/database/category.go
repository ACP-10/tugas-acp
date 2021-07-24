package database

import (
	"tugas-acp/configs"
	"tugas-acp/models/category"
)

func GetAllCategory() (dataResult []category.Category, err error) {
	err = configs.DB.Find(&dataResult).Error
	if err != nil {
		return nil, err
	}
	return
}
