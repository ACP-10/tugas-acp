package controllers

import (
	"net/http"
	"strconv"
	"tugas-acp/configs"
	"tugas-acp/lib/database"
	"tugas-acp/models/cart"
	cartitem "tugas-acp/models/cartItem"
	"tugas-acp/models/product"

	"github.com/labstack/echo/v4"
)

func CreateCartItemController(c echo.Context) error {
	var cartItemCreate cartitem.CartItemCreate
	c.Bind(&cartItemCreate)
	cartId, er := strconv.Atoi(c.Param("id"))

	if er != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"CartId not found",
			"",
		))
	}

	var cartItemDb cartitem.CartItem
	cartItemDb.CartId = cartId
	cartItemDb.ProductId = cartItemCreate.ProductId
	cartItemDb.Quantity = cartItemCreate.Quantity

	// update quantity to product table
	var productUpdate product.ProductUpdate
	productId := cartItemCreate.ProductId
	c.Bind(&productUpdate)

	var productDB product.Product
	configs.DB.First(&productDB, "product_id", productId)

	if productDB.Stock < cartItemCreate.Quantity {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Create Data",
			"Stock not available",
		))
	}

	productDB.Stock -= cartItemCreate.Quantity
	e := configs.DB.Save(&productDB).Error

	err := configs.DB.Create(&cartItemDb).Error

	if err != nil || e != nil {
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

	var cartDB cart.Cart
	er := configs.DB.First(&cartDB, "cart_id", id).Error

	if er != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"CartId not found",
			"",
		))
	}

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

	cartId, _ := strconv.Atoi(c.Param("cartId"))
	var cartDB cart.Cart
	er := configs.DB.First(&cartDB, "cart_id", cartId).Error

	if er != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"CartId not found",
			"",
		))
	}

	c.Bind(&cartItemUpdate)

	var cartItemDb cartitem.CartItem
	error := configs.DB.First(&cartItemDb, "item_id", cartItemId).Error

	if error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"CartItemId not found",
			"",
		))
	}

	//update product quantity
		var productDB product.Product
		configs.DB.First(&productDB, "product_id", cartItemDb.ProductId)

		if productDB.Stock + cartItemDb.Quantity < cartItemUpdate.Quantity {
			return c.JSON(http.StatusInternalServerError, BaseResponse(
				http.StatusInternalServerError,
				"Failed Create Data",
				"Stock not available",
			))
		}

		productDB.Stock += cartItemDb.Quantity - cartItemUpdate.Quantity
		e := configs.DB.Save(&productDB).Error

	cartItemDb.Quantity = cartItemUpdate.Quantity
	err := configs.DB.Save(&cartItemDb).Error

	if err != nil || e != nil{
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

func DeleteCartItemController(c echo.Context) error {
	cartItemId, _ := strconv.Atoi(c.Param("id"))

	cartId, _ := strconv.Atoi(c.Param("cartId"))
	var cartDB cart.Cart
	er := configs.DB.First(&cartDB, "cart_id", cartId).Error

	if er != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"CartId not found",
			"",
		))
	}

	var cartItemDB cartitem.CartItem

	e := configs.DB.First(&cartItemDB, "item_id", cartItemId).Error

	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"CartItemId not found",
			"",
		))
	}

	err := configs.DB.Where("item_id", cartItemId).Delete(&cartItemDB).Error

	var productDB product.Product
		configs.DB.First(&productDB, "product_id", cartItemDB.ProductId)
		productDB.Stock += cartItemDB.Quantity
		configs.DB.Save(&productDB)

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
