package controllers

import (
	"net/http"
	"strconv"
	"tugas-acp/configs"
	"tugas-acp/lib/database"
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

	return c.JSON(http.StatusCreated, categoryDB)
}

func GetCategoryController(c echo.Context) error {
	var categoryData []category.Category
	var err error

	categoryData, err = database.GetCategoryAll()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			categoryData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		categoryData,
	))
}

func UpdateCategoryController(c echo.Context) error {
	var categoryUpdate category.CategoryUpdate
	category_id, _ := strconv.Atoi(c.Param("id"))

	c.Bind(&categoryUpdate)

	var categoryDB category.Category
	configs.DB.First(&categoryDB, "category_id", category_id)
	categoryDB.CategoryName = categoryUpdate.CategoryName
	err := configs.DB.Save(&categoryDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, categoryDB)
}

func DeleteCategoryController(c echo.Context) error {
	category_id, _ := strconv.Atoi(c.Param("id"))

	var categoryDB category.Category
	err := configs.DB.Where("category_id", category_id).Delete(&categoryDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, categoryDB)
}
