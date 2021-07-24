package database

import (
	"tugas-acp/configs"
	"tugas-acp/models/product"
)

func GetAllProductByCategory(categoryId int) (dataResult []product.Product, err error) {
	err = configs.DB.Find(&dataResult, "category_id", categoryId).Error
	if err != nil {
		return nil, err
	}
	return
}
