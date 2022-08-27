package route

import (
	"github.com/osmanonurcan/go-test/api"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)

	e.GET("/plans", api.GetPlans)
	e.POST("/plan", api.AddPlan)
	e.PUT("/plan/:id", api.UpdatePlan)
	e.DELETE("/plan/:id", api.DeletePlan)

	return e
}
