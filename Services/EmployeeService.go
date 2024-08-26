package Services

import (
	"fmt"

	"github.com/gedehariyogananda/pattern-golang/Models"
	"github.com/gedehariyogananda/pattern-golang/Models/Common"
	"github.com/gedehariyogananda/pattern-golang/Models/Dto"
	"github.com/gedehariyogananda/pattern-golang/Repositories"
)

type (
	IEmployeeService interface {
		GetAllEmployees(perPage int, page int) (employees []Models.Employee, paginate Common.Meta, err error)
		GetEmployeeById(id string) (employee *Models.Employee, err error)
		AddNewEmployee(request *Dto.CreateEmployeeRequest) (err error)
		UpdateEmployee(id string, request *Dto.CreateEmployeeRequest) (employee *Models.Employee, err error)
		DeleteEmployee(id string) (err error)
	}

	EmployeeService struct {
		EmployeeRepository Repositories.IEmployeeRepository
	}
)

func EmployeeServiceProvider(employee Repositories.IEmployeeRepository) *EmployeeService {
	return &EmployeeService{
		EmployeeRepository: employee,
	}
}

func (h *EmployeeService) GetAllEmployees(perPage int, page int) (employees []Models.Employee, paginate Common.Meta, err error) {
	totalData, employees, err := h.EmployeeRepository.GetAll(perPage, page)

	if err != nil {
		return employees, paginate, err
	}

	// paginate.Data = employees
	paginate = Common.PaginateMetadata(totalData, perPage, page, "/employee")

	return employees, paginate, nil

}

func (h *EmployeeService) GetEmployeeById(id string) (employee *Models.Employee, err error) {
	employee, err = h.EmployeeRepository.FindId(id)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (h *EmployeeService) AddNewEmployee(request *Dto.CreateEmployeeRequest) (err error) {
	if _, err = h.EmployeeRepository.Create(request); err != nil {
		return fmt.Errorf("Error creating employee: %v", err)
	}

	return nil
}

func (h *EmployeeService) UpdateEmployee(id string, request *Dto.CreateEmployeeRequest) (employee *Models.Employee, err error) {
	employee, err = h.EmployeeRepository.FindId(id)

	if err != nil {
		return nil, err
	}

	if err = h.EmployeeRepository.Update(id, request); err != nil {
		return nil, err
	}

	return employee, nil

}

func (h *EmployeeService) DeleteEmployee(id string) (err error) {
	_, err = h.EmployeeRepository.FindId(id)

	if err != nil {
		return err
	}

	if err = h.EmployeeRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
