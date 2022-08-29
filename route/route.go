package route

import (
	"github.com/osmanonurcan/go-test/api"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)

	e.GET("/student/:id/plans", api.GetPlans)
	e.GET("/student/:id", api.GetStudent)
	e.GET("/students", api.GetStudents)

	//e.POST("/plan", api.AddPlan)
	e.POST("/student/:id/plan", api.AddPlan)
	e.POST("/student", api.AddStudent)

	e.PUT("student/:id", api.UpdateStudent)
	e.PUT("/student/:id/plan/:plan_id", api.UpdatePlan)

	e.DELETE("/student/:id", api.DeleteStudent)
	e.DELETE("/plan/:plan_id", api.DeletePlan)

	return e
}
