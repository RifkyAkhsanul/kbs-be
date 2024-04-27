package routes

import (
	"github.com/RifkyAkhsanul/kbs-be/controllers"
	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/info", controllers.GetInfo)
	e.POST("/infobyid", controllers.GetInfofromID)

	return e
}
