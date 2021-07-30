package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tugas-acp/models/product"

	"github.com/stretchr/testify/assert"
)


func TestGetProductByCategoryController(t *testing.T) {
	e := InitEcho()
	CreateSeedCategory()
	CreateSeedProduct()

	req,_ := http.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/product/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

		if assert.NoError(t, GetProductByCategoryController(c)) {

			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			baseResponse := product.ProductResponse{}
			if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
				assert.Error(t, err, "Failed convert body to object")
			}
			assert.Equal(t, http.StatusOK, baseResponse.Code)
		}
}

func TestCreateProductController(t *testing.T)  {
	e := InitEcho()

	input := []byte(
		`{
		"productName": "RedmiBook 15",
		"productPrice": 19999000,
		"stock": 100,
		"categoryId": 1
		}`)

	req,_ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(input))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/product")

	if assert.NoError(t, CreateProductController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		body := rec.Body.String()
		baseResponse := product.ProductResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}

func TestUpdateProductController(t *testing.T)  {
	e := InitEcho()

	CreateSeedCategory()
	CreateSeedProduct()

	input := []byte(
		`{
		"productName": "RedmiBook 15",
		"productPrice": 14999000,
		"stock": 50,
		"categoryId": 1
	}`)

	req,_ := http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(input))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/product/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, UpdateProductController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := product.ProductResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}

func TestDeleteProductController(t *testing.T)  {
	e := InitEcho()

	CreateSeedCategory()
	CreateSeedProduct()

	req,_ := http.NewRequest(http.MethodDelete, "/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/product/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteProductController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := product.ProductResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}