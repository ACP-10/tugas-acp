package controllers

import (
	"net/http"
	"tugas-acp/configs"
	"tugas-acp/middlewares"
	"tugas-acp/models/customer"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func RegisterCustomerController(c echo.Context) error {
	var customerRegister customer.CustomerCreate
	c.Bind(&customerRegister)

	hashedPassword, err := HashPassword(customerRegister.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var customerDB customer.Customer
	customerDB.Name = customerRegister.Name
	customerDB.Email = customerRegister.Email
	customerDB.Password = string(hashedPassword)

	e := configs.DB.Create(&customerDB).Error
	if e != nil {
		return c.JSON(http.StatusInternalServerError, e.Error())
	}

	var userResponse = customer.CustomerResponse{
		CustomerId: customerDB.CustomerId,
		Name:       customerDB.Name,
		Email:      customerDB.Email,
		CreatedAt:  customerDB.CreatedAt,
		UpdatedAt:  customerDB.UpdatedAt,
	}
	return c.JSON(http.StatusCreated, userResponse)
}

func LoginCustomerController(c echo.Context) error {
	customerLogin := customer.CustomerLogin{}
	c.Bind(&customerLogin)

	var customerDB customer.Customer
	configs.DB.First(&customerDB, "email", customerLogin.Email)
	hashedPassword := customerDB.Password

	err := VerifyPassword(hashedPassword, customerLogin.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	token, _ := middlewares.GenerateJWT(customerDB.CustomerId)

	var userResponse = customer.LoginResponse{
		Email: customerDB.Email,
		Token: token,
	}
	return c.JSON(http.StatusOK, userResponse)
}
