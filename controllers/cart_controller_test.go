package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"tugas-acp/configs"
	"tugas-acp/middlewares"
	"tugas-acp/models/cart"
	"tugas-acp/models/customer"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func InitEcho() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

func CreateSeedCart()  {
	var cartDB cart.Cart
	cartDB.IsCheckout = false
	cartDB.CustomerId = 1
	configs.DB.Create(&cartDB)
}

func CreateSeedCustomer()  {
	var customerDB customer.Customer
	customerDB.Name = "coba customer"
	customerDB.Email = "coba email customer"
	customerDB.Password = "$2a$10$haoCB7WNi.cWxPyymGeubed.BXxra4f2DwIGsKy3URVMIC/iOkqwO"
	configs.DB.Create(&customerDB)
}

//error
func TestGetCartController(t *testing.T)  {
	e := InitEcho()
	CreateSeedCustomer()
	CreateSeedCart()
	token,_ := middlewares.GenerateJWT(1)

	// auth := "Bearer " + token
	// fmt.Println("\n" + string(auth) + "\n")

	req,_ := http.NewRequest(http.MethodGet, "/cart", nil)
	
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(c)
	// jwt := middleware.JWTWithConfig(c)
	// fmt.Println(&jwt)
	// c.SetPath("/cart")


	if assert.NoError(t, GetCartController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		fmt.Println(rec.Body.String())
		body := rec.Body.String()
		baseResponse := cart.CartResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}

func TestCreateCartController(t *testing.T)  {
	
}