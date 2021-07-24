package controllers

import (
	"net/http"
	"strconv"
	"tugas-acp/configs"
	"tugas-acp/lib/database"
	"tugas-acp/middlewares"
	"tugas-acp/models/cart"
	cartitem "tugas-acp/models/cartItem"

	"github.com/labstack/echo/v4"
)


func CreateCartItemController(c echo.Context) error{
	var cartItemCreate cartitem.CartItemCreate
	var cartCheck cart.Cart
	var cartId int
	c.Bind(&cartItemCreate)

	er := configs.DB.First(&cartCheck, "cart_id", cartItemCreate.CartId).Error

	if er != nil {

		customerId := middlewares.GetUserIdFromJWT(c)

		var cartDB cart.Cart
		cartDB.IsCheckout = false
		cartDB.CustomerId = customerId

		err := configs.DB.Create(&cartDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

		// return c.JSON(http.StatusInternalServerError, BaseResponse(
		// 	http.StatusInternalServerError,
		// 	"CartId not found",
		// 	"",
		// ))
		cartId = cartDB.CartId
	}

	var cartItemDb cartitem.CartItem
	if cartId != 0 {
		cartItemDb.CartId = cartId
	}else{
		cartItemDb.CartId = cartItemCreate.CartId
	}
	cartItemDb.ProductId = cartItemCreate.ProductId
	cartItemDb.Quantity = cartItemCreate.Quantity

	err := configs.DB.Create(&cartItemDb).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Create Data",
			err.Error(),
		))
	}

	if cartId != 0 {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusOK,
			"Success Create Data Cart&CartItem",
			cartItemDb,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Create Data CartItem",
		cartItemDb,
	))
}

func GetCartItemControllers(c echo.Context) error {
	var cartItemData []cartitem.CartItemResult
	var err error

	id, _ := strconv.Atoi(c.Param("id"))

	cartItemData, err = database.GetCartItemByCartId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			cartItemData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		cartItemData,
	))
}

func UpdateCartItemController(c echo.Context) error {
	var cartItemUpdate cartitem.CartItemUpdate
	cartItemId, _ := strconv.Atoi(c.Param("id"))

	c.Bind(&cartItemUpdate)

	var cartItemDb cartitem.CartItem
	configs.DB.First(&cartItemDb, "item_id", cartItemId)
	cartItemDb.Quantity = cartItemUpdate.Quantity
	err := configs.DB.Save(&cartItemDb).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Update Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Update Data CartItem",
		cartItemDb,
	))
}

func DeleteCartItemController(c echo.Context)error{
	cartItemId, _ := strconv.Atoi(c.Param("id"))

	var cartItemDB cartitem.CartItem
	err := configs.DB.Where("item_id", cartItemId).Delete(&cartItemDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Delete Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success delete CartItem",
		"",
	))
}