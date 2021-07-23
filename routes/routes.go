package routes

import (
	"tugas-acp/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/cart", controllers.GetCartController)
	e.POST("/cart", controllers.CreateCartController)
	e.POST("/category", controllers.CreateCategoryController)

	return e
}
