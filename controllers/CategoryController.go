package controllers

import (
	"net/http"
	"tugas-acp/configs"
	"tugas-acp/models/category"

	"github.com/labstack/echo/v4"
)

func CreateCategoryController(c echo.Context) error {
	var categoryCreate category.CategoryCreate
	c.Bind(&categoryCreate)

	var categoryDB category.Category
	categoryDB.CategoryName = categoryCreate.CategoryName

	err := configs.DB.Create(&categoryDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, categoryDB)
}
