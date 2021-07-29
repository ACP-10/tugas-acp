package database

import (
	"tugas-acp/configs"
	"tugas-acp/models/product"
)

func GetProductByCategoryId(categoryId int) (dataresult []product.ProductResult, err error) {
	err = configs.DB.Model(&product.Product{}).
		Select("categories.category_name AS category, products.product_id, products.product_name, products.product_price, products.stock").
		Joins("JOIN categories ON categories.category_id = products.category_id").
		Where("products.category_id = ?", categoryId).Scan(&dataresult).Error

	if err != nil {
		return nil, err
	}
	return
}
