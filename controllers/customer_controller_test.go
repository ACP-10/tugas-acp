package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tugas-acp/models/customer"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func TestRegisterCustomer(t *testing.T) {
	e := InitEcho()

	input := []byte(`
	{"name": "testing",
    "email": "testingCustomer@gmail.com",
    "password": "Password123"}`)

	req,_ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(input))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/register")

	if assert.NoError(t, RegisterCustomerController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		body := rec.Body.String()
		baseResponse := customer.CustomerResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		// assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}

// error
func TestLoginCustomer(t *testing.T)  {
	e := InitEcho()
	CreateSeedCustomer()

	input := []byte(`{"email": "testingCustomer@gmail.com","password": "Password123"}`)

	req,_ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(input))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/customer/login")

	if assert.NoError(t, LoginCustomerController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := customer.CustomerResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
	}
}