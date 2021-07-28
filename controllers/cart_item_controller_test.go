package controllers

import (
	// "encoding/json"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	cartitem "tugas-acp/models/cartItem"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCartItemController(t *testing.T) {
	e := InitEcho()
	CreateSeedCustomer()
	CreateSeedProduct()
	CreateSeedCategory()
	CreateSeedCart()
	CreateSeedCartItem()

	req,_ := http.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart/:id/item")
	c.SetParamNames("id")
	c.SetParamValues("1")

		if assert.NoError(t, GetCartItemControllers(c)) {

			assert.Equal(t, http.StatusOK, rec.Code)
			// fmt.Println(rec.Body.String())
			body := rec.Body.String()
			baseResponse := cartitem.CartItemResponse{}
			if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
				assert.Error(t, err, "Failed convert body to object")
			}
			assert.Equal(t, http.StatusOK, baseResponse.Code)
		}

}

func TestCreateCartItemController(t *testing.T)  {
	e := InitEcho()

	input := []byte(`{"productId": 1, "quantity" : 10}`)

	req,_ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(input))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart/:id/item")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, CreateCartItemController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := cartitem.CartItemResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		t.Log(baseResponse)
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}

func TestUpdateCartItemController(t *testing.T)  {
	e := InitEcho()

	CreateSeedCart()
	CreateSeedCartItem()

	input := []byte(`{"quantity" : 10}`)

	req,_ := http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(input))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart/:cartId/item/:id")
	c.SetParamNames("cartId","id")
	c.SetParamValues("1","1")

	if assert.NoError(t, UpdateCartItemController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := cartitem.CartItemResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}

func TestDeleteCartItemController(t *testing.T)  {
	e := InitEcho()

	CreateSeedCart()
	CreateSeedCartItem()

	req,_ := http.NewRequest(http.MethodDelete, "/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart/:cartId/item/:id")
	c.SetParamNames("cartId","id")
	c.SetParamValues("1","1")

	if assert.NoError(t, DeleteCartItemController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := cartitem.CartItemResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}