package Routes

import (
	"github.com/gedehariyogananda/pattern-golang/Routes/Di"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EmployeeRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/employees")

	// the route access must be authenticated
	m := Di.DICommonMiddleware(db)

	// open use access authenticate
	route.Use(m.IsAuthenticate)

	// init controller
	EmployeeController := Di.DIEmployee(db)

	route.GET("/checked", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "list division",
		})
	})

	route.GET("/", EmployeeController.GetAllEmployees)
	route.GET("/:id", EmployeeController.GetEmployeeById)
	route.POST("/", EmployeeController.AddNewEmployee)
	route.PUT("/:id", EmployeeController.UpdateEmployee)
	route.DELETE("/:id", EmployeeController.DeleteEmployee)
}
