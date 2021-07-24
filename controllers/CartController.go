package controllers

import (
	"net/http"
	"strconv"
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
		return c.JSON(http.StatusInternalServerError, BaseResponse(
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

func UpdateCartController(c echo.Context) error {
	var cartUpdate cart.CartUpdate
	cart_id, _ := strconv.Atoi(c.Param("id"))

	c.Bind(&cartUpdate)

	var cartDB cart.Cart
	configs.DB.First(&cartDB, "cart_id", cart_id)
	cartDB.IsCheckout = cartUpdate.IsCheckout
	err := configs.DB.Save(&cartDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, cartDB)
}

func DeleteCartController(c echo.Context) error {
	cart_id, _ := strconv.Atoi(c.Param("id"))

	var cartDB cart.Cart
	err := configs.DB.Where("cart_id", cart_id).Delete(&cartDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Success delete Cart")
}