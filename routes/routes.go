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

	return e
}
