package routes

import (
	"tugas-acp/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/category", controllers.GetCategoryController)
	e.POST("/category", controllers.CreateCategoryController)
	e.PUT("/category/:id", controllers.UpdateCategoryController)
	e.DELETE("/category/:id", controllers.DeleteCategoryController)

	e.GET("/cart", controllers.GetCartController)
	e.POST("/cart", controllers.CreateCartController)
	e.PUT("/cart/:id", controllers.UpdateCartController)
	e.DELETE("/cart/:id", controllers.DeleteCartController)

	e.POST("/customer/register", controllers.RegisterCustomerController)
	e.POST("/customer/login", controllers.LoginCustomerController)

	e.GET("/cartItem/:id", controllers.GetCartItemControllers)
	e.POST("/cartItem", controllers.CreateCartItemController)
	e.PUT("/cartItem/:id", controllers.UpdateCartItemController)
	e.DELETE("/cartItem/:id", controllers.DeleteCartItemController)

	return e
}
