package Routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(c *gin.RouterGroup, db *gorm.DB) {
	route := c.Group("/auth")

	route.GET("/checked", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "checked healt",
		})
	})
}
