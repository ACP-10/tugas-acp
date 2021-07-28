package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tugas-acp/configs"
	cartitem "tugas-acp/models/cartItem"
	"tugas-acp/models/category"
	"tugas-acp/models/product"

	"github.com/stretchr/testify/assert"
)

func CreateSeedProduct()  {
	var ProductDB product.Product
	ProductDB.ProductName = "PS4 Slim"
	ProductDB.CategoryId = 1
	ProductDB.ProductPrice = 2500000
	ProductDB.Stock = 100
	configs.DB.Create(ProductDB)
}

func CreateSeedCartItem()  {
	var cartItemDB cartitem.CartItem
	cartItemDB.CartId = 1
	cartItemDB.ProductId = 1
	cartItemDB.Quantity = 10
	configs.DB.Create(cartItemDB)
}

func CreateSeedCategory()  {
	var categoryDB category.Category
	categoryDB.CategoryName = "Konsol"
}

func TestGetCartItemController(t *testing.T) {
	e := InitEcho()
	CreateSeedCustomer()
	CreateSeedCart()
	req,_ := http.NewRequest(http.MethodGet, "/cart/1/item", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetCartItemControllers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// fmt.Println(rec.Body.String())
		body := rec.Body.String()
		// baseResponse := cart.CartResponse{}
		baseResponse := cartitem.CartItemResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}