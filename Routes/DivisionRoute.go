package Routes

import (
	"github.com/gedehariyogananda/pattern-golang/Routes/Di"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DivisionRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/division")

	// the route access must be authenticated
	m := Di.DICommonMiddleware(db)

	// open use access authenticate
	route.Use(m.IsAuthenticate)

	// init controller
	divisionController := Di.DIDivision(db)

	route.GET("/checked", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "list division",
		})
	})

	route.GET("/", divisionController.GetAllDivision)
}
