package controllers

import (
	"net/http"
	"tugas-acp/configs"
	"tugas-acp/lib/database"
	"tugas-acp/models/cart"

	"github.com/labstack/echo/v4"
)

func CreateCartController(c echo.Context) error {

	var cartCreate cart.CartCreate
	c.Bind(&cartCreate)

	var cartDB cart.Cart
	cartDB.IsCheckout = cartCreate.IsCheckout
	cartDB.CustomerId = cartCreate.CustomerId

	err := configs.DB.Create(&cartDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartDB)
}

func GetCartController(c echo.Context) error {
	var cartData []cart.Cart
	var err error

	cartData, err = database.GetCartAll()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			cartData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data Cart",
		cartData,
	))
}

// func UpdateCartController(c echo.Context) error {
	
// }

// func DeleteCartController(c echo.Context) error {
	
// }