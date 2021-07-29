package routes

import (
	"tugas-acp/constants"
	"tugas-acp/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET)))

	e.GET("/test/category", controllers.GetCategoryController)
	e.POST("/test/category", controllers.CreateCategoryController)
	e.PUT("/test/category/:id", controllers.UpdateCategoryController)
	e.DELETE("/test/category/:id", controllers.DeleteCategoryController)

	eAuth.GET("/category", controllers.GetCategoryController)
	eAuth.POST("/category", controllers.CreateCategoryController)
	eAuth.PUT("/category/:id", controllers.UpdateCategoryController)
	eAuth.DELETE("/category/:id", controllers.DeleteCategoryController)

	eAuth.GET("/cart", controllers.GetCartController)
	eAuth.POST("/cart", controllers.CreateCartController)
	eAuth.PUT("/cart/:id", controllers.UpdateCartController)
	eAuth.DELETE("/cart/:id", controllers.DeleteCartController)

	e.POST("/customer/register", controllers.RegisterCustomerController)
	e.POST("/customer/login", controllers.LoginCustomerController)

	eAuth.GET("/cartItem/:id", controllers.GetCartItemControllers)
	eAuth.POST("/cartItem", controllers.CreateCartItemController)
	eAuth.PUT("/cartItem/:id", controllers.UpdateCartItemController)
	eAuth.DELETE("/cartItem/:id", controllers.DeleteCartItemController)

	eAuth.GET("/product/:id", controllers.GetProductByCategoryController)
	eAuth.POST("/product", controllers.CreateProductControler)
	eAuth.PUT("/product/:id", controllers.UpdateProductController)
	eAuth.DELETE("/product/:id", controllers.DeleteProductController)

	eAuth.POST("/payment/:id", controllers.Pay)

	return e
}
