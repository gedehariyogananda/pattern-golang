package Dto

type CreateEmployeeRequest struct {
	Name       string `form:"name" binding:"required"`
	Phone      string `form:"phone" binding:"required"`
	Position   string `form:"position" binding:"required"`
	DivisionId string `form:"division_id" binding:"required"`
	Image      string
}
