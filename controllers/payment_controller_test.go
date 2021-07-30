package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tugas-acp/models/payment"

	"github.com/stretchr/testify/assert"
)


func TestPay(t *testing.T) {
	e := InitEcho()

	CreateSeedCartItem()
	CreateSeedCart()

	req,_ := http.NewRequest(http.MethodPost, "/",nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/payment/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, Pay(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := payment.PaymentResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}