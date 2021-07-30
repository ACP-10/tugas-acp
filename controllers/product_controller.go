package controllers

import (
	"net/http"
	"strconv"
	"tugas-acp/configs"
	"tugas-acp/lib/database"
	"tugas-acp/models/product"

	"github.com/labstack/echo/v4"
)

func CreateProductController(c echo.Context) error {
	var productCreate product.ProductCreate
	c.Bind(&productCreate)

	var productDB product.Product
	productDB.ProductName = productCreate.ProductName
	productDB.ProductPrice = productCreate.ProductPrice
	productDB.Stock = productCreate.Stock
	productDB.CategoryId = productCreate.CategoryId

	err := configs.DB.Create(&productDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Create Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusOK,
		"Success Create Data Product",
		productDB,
	))
}

func GetProductByCategoryController(c echo.Context) error {
	var productData []product.ProductResult
	var err error
	categoryId, _ := strconv.Atoi(c.Param("id"))

	productData, err = database.GetProductByCategoryId(categoryId)

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			productData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		productData,
	))
}

func UpdateProductController(c echo.Context) error {
	var productUpdate product.ProductUpdate
	productId, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&productUpdate)

	var productDB product.Product
	configs.DB.First(&productDB, "product_id", productId)
	productDB.ProductName = productUpdate.ProductName
	productDB.ProductPrice = productUpdate.ProductPrice
	productDB.Stock = productUpdate.Stock
	productDB.CategoryId = productUpdate.CategoryId
	err := configs.DB.Save(&productDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Update Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Update Data Product",
		productDB,
	))
}

func DeleteProductController(c echo.Context) error {
	productId, _ := strconv.Atoi(c.Param("id"))
	var productDB product.Product

	err := configs.DB.Where("product_id", productId).Delete(&productDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Delete Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success delete Product",
		"",
	))
}
