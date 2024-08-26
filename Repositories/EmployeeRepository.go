package Repositories

import (
	"errors"
	"fmt"

	"github.com/gedehariyogananda/pattern-golang/Models"
	"github.com/gedehariyogananda/pattern-golang/Models/Dto"
	"github.com/gedehariyogananda/pattern-golang/Utils"
	"gorm.io/gorm"
)

type (
	IEmployeeRepository interface {
		GetAll(perPage int, page int) (totalData int64, employees []Models.Employee, err error)
		FindId(id string) (employees *Models.Employee, err error)
		Create(request *Dto.CreateEmployeeRequest) (employee *Models.Employee, err error)
		Update(id string, request *Dto.CreateEmployeeRequest) (err error)
		Delete(id string) (err error)
	}

	EmployeeRepository struct {
		DB *gorm.DB
	}
)

func EmployeeRepositoryProvider(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (h *EmployeeRepository) GetAll(perPage int, page int) (totalData int64, employees []Models.Employee, err error) {

	if err := h.DB.Model(&Models.Employee{}).Count(&totalData).Error; err != nil {
		return 0, nil, err
	}

	dataset := h.DB.Scopes(Utils.Paginate(page, perPage)).Preload("Division").Find(&employees)

	if dataset.Error != nil {
		return 0, nil, dataset.Error
	}

	if dataset.RowsAffected == 0 {
		return 0, nil, errors.New("data not found")
	}

	return totalData, employees, nil
}

func (h *EmployeeRepository) FindId(id string) (employees *Models.Employee, err error) {
	if err = h.DB.Preload("Division").Where("id = ?", id).First(&employees).Error; err != nil {
		return nil, errors.New("data not found")
	}

	return employees, nil
}

func (h *EmployeeRepository) Create(request *Dto.CreateEmployeeRequest) (employees *Models.Employee, err error) {
	tx := h.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	employees = &Models.Employee{
		Name:       request.Name,
		Phone:      request.Phone,
		Image:      request.Image,
		Position:   request.Position,
		DivisionId: request.DivisionId,
	}

	if err := h.DB.Create(&employees); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error create employee: %v", err.Error)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("error commit transaction: %v", err.Error)
	}

	return employees, nil
}

func (h *EmployeeRepository) Update(id string, request *Dto.CreateEmployeeRequest) (err error) {

	employee := &Models.Employee{
		Name:       request.Name,
		Phone:      request.Phone,
		Image:      request.Image,
		Position:   request.Position,
		DivisionId: request.DivisionId,
	}

	if err := h.DB.Model(&Models.Employee{}).Where("id = ?", id).Updates(&employee); err != nil {
		return err.Error
	}

	return nil
}

func (h *EmployeeRepository) Delete(id string) (err error) {
	if err := h.DB.Where("id = ?", id).Delete(&Models.Employee{}); err != nil {
		return errors.New("data not found")
	}

	return nil
}
