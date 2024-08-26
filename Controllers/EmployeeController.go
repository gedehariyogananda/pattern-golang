package Controllers

import (
	"github.com/gedehariyogananda/pattern-golang/Models/Dto"
	"github.com/gedehariyogananda/pattern-golang/Services"
	"github.com/gedehariyogananda/pattern-golang/Utils"
	"github.com/gin-gonic/gin"
)

type (
	IEmployeeController interface {
		GetAllEmployees(ctx *gin.Context)
		GetEmployeeById(ctx *gin.Context)
		AddNewEmployee(ctx *gin.Context)
		UpdateEmployee(ctx *gin.Context)
		DeleteEmployee(ctx *gin.Context)
	}

	EmployeeController struct {
		EmployeeService Services.IEmployeeService
	}
)

func EmployeeControllerProvider(employeeService Services.IEmployeeService) *EmployeeController {
	return &EmployeeController{EmployeeService: employeeService}
}

func (c *EmployeeController) GetAllEmployees(ctx *gin.Context) {

	// setup pagination
	defaultPerPage := 3
	defaultPage := 1

	perPage, page := Utils.GetPaginationParams(ctx, defaultPerPage, defaultPage)

	employees, employeesPaginate, err := c.EmployeeService.GetAllEmployees(perPage, page)

	if err != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "mantap",
		"data": gin.H{
			"employees": employees,
		},
		"meta": employeesPaginate,
	})
}

func (c *EmployeeController) GetEmployeeById(ctx *gin.Context) {
	id := ctx.Param("id")
	employee, err := c.EmployeeService.GetEmployeeById(id)

	if err != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "mantap",
		"data":    employee,
	})
}

func (c *EmployeeController) AddNewEmployee(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	var request Dto.CreateEmployeeRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": "Failed to bind request: " + err.Error(),
		})
		return
	}

	fileName := Utils.GenerateUniqueFileName(fileHeader.Filename)
	request.Image = "/uploads/" + fileName

	if err := c.EmployeeService.AddNewEmployee(&request); err != nil {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	filePath := "./public/uploads/" + fileName
	successUpload := ctx.SaveUploadedFile(fileHeader, filePath)

	if successUpload != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": successUpload.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "mantap",
	})

}

func (c *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	id := ctx.Param("id")
	var request *Dto.CreateEmployeeRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, err := c.EmployeeService.UpdateEmployee(id, request)

	if err != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "mantap",
		"data":    user,
	})
}

func (c *EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.EmployeeService.DeleteEmployee(id)

	if err != nil {
		ctx.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "mantap",
	})
}
