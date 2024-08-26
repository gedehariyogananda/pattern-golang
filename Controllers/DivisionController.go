package Controllers

import (
	"net/http"

	"github.com/gedehariyogananda/pattern-golang/Services"
	"github.com/gedehariyogananda/pattern-golang/Utils"
	"github.com/gin-gonic/gin"
)

type (
	IDivisionController interface {
		GetAllDivision(ctx *gin.Context)
	}

	DivisionController struct {
		DivisionService Services.IDivisionService
	}
)

func DivisionControllerPrivoder(divisionService Services.IDivisionService) *DivisionController {
	return &DivisionController{DivisionService: divisionService}
}

func (c *DivisionController) GetAllDivision(ctx *gin.Context) {

	// setup pagination
	defaultPerPage := 3
	defaultPage := 1

	perPage, page := Utils.GetPaginationParams(ctx, defaultPerPage, defaultPage)

	divisions, divisionPaginate, err := c.DivisionService.GetAllDivison(perPage, page)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "mantap",
		"data":  divisions,
		"meta": divisionPaginate,	
	})

}
