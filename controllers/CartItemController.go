package controllers

import (
	"net/http"
	"strconv"
	"tugas-acp/configs"
	"tugas-acp/lib/database"
	cartitem "tugas-acp/models/cartItem"

	"github.com/labstack/echo/v4"
)


func CreateCartItemController(c echo.Context) error{
	var cartItemCreate cartitem.CartItemCreate
	c.Bind(&cartItemCreate)

	var cartItemDb cartitem.CartItem
	cartItemDb.CartId = cartItemCreate.CartId
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