package controllers

import (
	"net/http"
	"strconv"
	"tugas-acp/configs"
	"tugas-acp/models/payment"

	"github.com/labstack/echo/v4"
)

func Pay(c echo.Context) error {
	var paymentUpdate payment.PaymentUpdate
	cartId, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&paymentUpdate)

	var paymentDB payment.Payment
	configs.DB.First(&paymentUpdate, "cart_id", cartId)
	paymentDB.IsPaid = true
	err := configs.DB.Save(&paymentDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Pay Transaction",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Pay Transaction",
		paymentDB,
	))
}
