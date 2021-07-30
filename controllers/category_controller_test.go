package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tugas-acp/models/category"

	"github.com/stretchr/testify/assert"
)


func TestGetCategoryController(t *testing.T) {
	e := InitEcho()
	CreateSeedCategory()

	req,_ := http.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/category")

		if assert.NoError(t, GetCategoryController(c)) {

			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			baseResponse := category.CategoryResponse{}
			if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
				assert.Error(t, err, "Failed convert body to object")
			}
			assert.Equal(t, http.StatusOK, baseResponse.Code)
		}
}

func TestCreateCategoryController(t *testing.T)  {
	e := InitEcho()

	input := []byte(`{"categoryName" : "laptop"}`)

	req,_ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(input))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/category")

	if assert.NoError(t, CreateCategoryController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		body := rec.Body.String()
		baseResponse := category.CategoryResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}

func TestUpdateCategoryController(t *testing.T)  {
	e := InitEcho()

	CreateSeedCategory()

	input := []byte(`{"categoryName": "smartwatch"}`)

	req,_ := http.NewRequest(http.MethodPut, "/", bytes.NewBuffer(input))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/category/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, UpdateCategoryController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := category.CategoryResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}

func TestDeleteCategoryController(t *testing.T)  {
	e := InitEcho()

	CreateSeedCategory()

	req,_ := http.NewRequest(http.MethodDelete, "/", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/category/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteCategoryController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()
		baseResponse := category.CategoryResponse{}
		if err := json.Unmarshal([]byte(body), &baseResponse); err != nil {
			assert.Error(t, err, "Failed convert body to object")
		}
		assert.Equal(t, http.StatusOK, baseResponse.Code)
	}
}