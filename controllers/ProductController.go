package controllers

import (
	"net/http"
	"strconv"
	"tugas-acp/configs"
	"tugas-acp/lib/database"
	"tugas-acp/models/product"

	"github.com/labstack/echo/v4"
)

func CreateProductControler(c echo.Context) error {
	var productCreate product.ProductCreate
	c.Bind(&productCreate)

	var productDB product.Product
	productDB.ProductName = productCreate.ProductName
	productDB.ProductPrice = productCreate.ProductPrice
	productDB.Stock = productCreate.Stock
	productDB.CategoryId = productCreate.CategoryId

	err := configs.DB.Create(&productDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, productDB)
}

func GetProductByCategoryController(c echo.Context) error {
	var productData []product.Product
	var err error
	categoryId, _ := strconv.Atoi(c.Param("categoryId"))

	productData, err = database.GetAllProductByCategory(categoryId)

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
