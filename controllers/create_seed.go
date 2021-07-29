package controllers

import (
	"tugas-acp/configs"
	"tugas-acp/models/cart"
	cartitem "tugas-acp/models/cartItem"
	"tugas-acp/models/category"
	"tugas-acp/models/customer"
	"tugas-acp/models/product"
)


func CreateSeedCart() {
	var cartDB cart.Cart
	cartDB.IsCheckout = false
	cartDB.CustomerId = 1
	configs.DB.Create(&cartDB)
}

func CreateSeedCustomer() {
	var customerDB customer.Customer
	customerDB.Name = "coba customer"
	customerDB.Email = "coba email customer"
	customerDB.Password = "$2a$10$haoCB7WNi.cWxPyymGeubed.BXxra4f2DwIGsKy3URVMIC/iOkqwO"
	configs.DB.Create(&customerDB)
}

func CreateSeedProduct()  {
	var ProductDB product.Product
	ProductDB.ProductName = "PS4 Slim"
	ProductDB.CategoryId = 1
	ProductDB.ProductPrice = 2500000
	ProductDB.Stock = 100
	configs.DB.Create(&ProductDB)
}

func CreateSeedCartItem()  {
	var cartItemDB cartitem.CartItem
	cartItemDB.CartId = 1
	cartItemDB.ProductId = 1
	cartItemDB.Quantity = 15
	configs.DB.Create(&cartItemDB)
}

func CreateSeedCategory()  {
	var categoryDB category.Category
	categoryDB.CategoryName = "Konsol"
	configs.DB.Create(&categoryDB)
}