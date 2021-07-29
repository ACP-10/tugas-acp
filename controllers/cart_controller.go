package controllers

import (
	"net/http"
	"strconv"
	"tugas-acp/configs"
	"tugas-acp/lib/database"
	"tugas-acp/middlewares"
	"tugas-acp/models/cart"
	cartitem "tugas-acp/models/cartItem"
	"tugas-acp/models/payment"

	"github.com/labstack/echo/v4"
)

func CreateCartController(c echo.Context) error {

	var cartCreate cart.CartCreate
	c.Bind(&cartCreate)

	customerId := middlewares.GetUserIdFromJWT(c)

	var cartDB cart.Cart
	cartDB.IsCheckout = cartCreate.IsCheckout
	cartDB.CustomerId = customerId

	err := configs.DB.Create(&cartDB).Error

	if err != nil {
		// return c.JSON(http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Create Cart",
			err.Error(),
		))
	}
	
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Create Cart",
		cartDB,
	))
}

func GetCartController(c echo.Context) error {
	var cartData cart.CartResult
	var err error
	var cartItem []cartitem.CartItemResult

	// jika ingin memasukkan customerId sendiri
	// customerId, _ := strconv.Atoi(c.QueryParam("customerId"))



	customerId := middlewares.GetUserIdFromJWT(c)

	cartData, err = database.GetCartByCustomer(customerId,false)
	cartItem, _ = database.GetCartItemByCartId(cartData.CartId)
	
	cartData.Item = cartItem

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			err.Error(),
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
	er := configs.DB.First(&cartDB, "cart_id", cart_id).Error

	if er != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"CartId not found",
			"",
		))
	}

	cartDB.IsCheckout = cartUpdate.IsCheckout
	err := configs.DB.Save(&cartDB).Error

	if cartUpdate.IsCheckout {
		// insert ke payment table
		var paymentCreate payment.PaymentCreate
		c.Bind(&paymentCreate)
		var paymentDB payment.Payment
		paymentDB.CartId = cart_id
		paymentDB.IsPaid = false
		total, e := database.GetTotalPayment(cart_id)
		if e != nil {
			// return c.JSON(http.StatusInternalServerError, e.Error())
			return c.JSON(http.StatusInternalServerError, BaseResponse(
				http.StatusInternalServerError,
				"Failed Get Total Payment",
				e.Error(),
			))
		}

		paymentDB.TotalPayment = total
		er := configs.DB.Create(&paymentDB).Error
		if er != nil {
			// return c.JSON(http.StatusInternalServerError, er.Error())

			return c.JSON(http.StatusInternalServerError, BaseResponse(
				http.StatusInternalServerError,
				"Failed Create Payment",
				er.Error(),
			))
		}

		// return c.JSON(http.StatusCreated, paymentDB)

		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusCreated,
			"Success Update Cart & Create Payment",
			paymentDB,
		))
	}

	if err != nil {
		// return c.JSON(http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Update Cart",
			err.Error(),
		))
	}

	// return c.JSON(http.StatusCreated, cartDB)

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Update Cart",
		cartDB,
	))
}

func DeleteCartController(c echo.Context) error {
	cart_id, _ := strconv.Atoi(c.Param("id"))

	var cartDB cart.Cart

	er := configs.DB.First(&cartDB, "cart_id", cart_id).Error

	if er != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusNotFound,
			"CartId not found",
			"",
		))
	}

	err := configs.DB.Where("cart_id", cart_id).Delete(&cartDB).Error
	configs.DB.Where("cart_id", cart_id).Delete(&cartitem.CartItem{})

	if err != nil {
		// return c.JSON(http.StatusInternalServerError, err.Error())
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Delete Cart",
			err.Error(),
		))
		
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success delete Cart",
		"",
	))
}
