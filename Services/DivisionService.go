package Services

import (
	// "os"
	// "strconv"

	"github.com/gedehariyogananda/pattern-golang/Models"
	"github.com/gedehariyogananda/pattern-golang/Models/Common"
	"github.com/gedehariyogananda/pattern-golang/Repositories"
)

type (
	IDivisionService interface {
		GetAllDivison(perPage int, page int) (divisions []Models.Division, paginate Common.Meta, err error)
	}

	DivisionService struct {
		DivisionRepository Repositories.IDivisionRepository
	}
)

func DivisionServiceProvider(productRepository Repositories.IDivisionRepository) *DivisionService {
	return &DivisionService{
		DivisionRepository: productRepository,
	}
}

func (h *DivisionService) GetAllDivison(perPage int, page int) (divisions []Models.Division, paginate Common.Meta, err error) {
	totalData, divisions, err := h.DivisionRepository.GetAll(perPage, page)

	if err != nil {
		return divisions, paginate, err
	}

	// paginate.Data = divisions
	paginate = Common.PaginateMetadata(totalData, perPage, page, "/division")

	return divisions, paginate, nil

}
