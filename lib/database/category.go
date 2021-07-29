package database

import (
	"tugas-acp/configs"
	"tugas-acp/models/category"
)

func GetCategoryAll() (dataresult []category.Category, err error) {
	err = configs.DB.Find(&dataresult).Error
	if err != nil {
		return nil, err
	}
	return
}
