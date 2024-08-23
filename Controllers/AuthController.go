package Controllers

import (
	"net/http"

	"github.com/gedehariyogananda/pattern-golang/Models/Dto"
	"github.com/gedehariyogananda/pattern-golang/Services"
	"github.com/gin-gonic/gin"
)

type (
	IAuthController interface {
		Register(ctx *gin.Context)
	}

	AuthController struct {
		service Services.IAuthService
	}
)

func AuthControllerProvider(service Services.IAuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var registerRequest Dto.RegisterRequest

	if err := ctx.ShouldBind(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := c.service.Register(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "mantap",
	})
}
