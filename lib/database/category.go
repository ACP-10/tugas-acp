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

func UpdateCategory(category_id int, category_name string) (dataresult []category.Category, err error) {
	err = configs.DB.First(&dataresult, "category_id = ?", category_id).Update("category_name", category_name).Error
	if err != nil {
		return nil, err
	}
	return
}
