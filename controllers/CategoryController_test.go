package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"tugas-acp/configs"
	"tugas-acp/models/category"

	"github.com/labstack/echo/v4"
)

func InitEcho() *echo.Echo {
	configs.InitConfigDBTest()
	e := echo.New()
	return e
}

func CreateSeedCategory() {
	var categoryDB category.Category
	categoryDB.CategoryName = "laptop"
	configs.DB.Create(&categoryDB)
}

func TestGetCategoryController(t *testing.T) {
	e := InitEcho()

	// CreateSeedCategory()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/category")

	// if assert.NoError(t, GetCategoryController(c)) {

	// 	assert.Equal(t, http.StatusOK, rec.Code)
	// 	body := rec.Body.String()
	// 	baseResponse := category.CategoryResponse{}
	// 	if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
	// 		assert.Error(t, err, "Failed convert body to object")
	// 	}
	// 	assert.Equal(t, http.StatusOK, baseResponse.Code)
	// }
}

func TestCreateCategoryController(t *testing.T) {
	e := InitEcho()
	input := []byte(`{"categoryName": "smartphone"}`)
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(input))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/test/category")

	// if assert.NoError(t, CreateCategoryController(c)) {
	// 	assert.Equal(t, http.StatusOK, rec.Code)
	// 	body := rec.Body.String()
	// 	baseResponse := category.CategoryResponse{}
	// 	if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
	// 		assert.Error(t, err, "Failed convert body to object")
	// 	}
	// 	t.Log(baseResponse)
	// 	assert.Equal(t, http.StatusOK, baseResponse.Code)
	// }
}
