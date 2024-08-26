package Repositories

import (
	"github.com/gedehariyogananda/pattern-golang/Models"
	"github.com/gedehariyogananda/pattern-golang/Utils"
	"gorm.io/gorm"
)

type (
	IDivisionRepository interface {
		GetAll(perPage int, page int) (totalData int64, divisions []Models.Division, err error)
	}

	DivisionRepository struct {
		DB *gorm.DB
	}
)

func DivisionRepositoryProvider(db *gorm.DB) *DivisionRepository {
	return &DivisionRepository{DB: db}
}

func (h *DivisionRepository) GetAll(perPage int, page int) (totalData int64, divisions []Models.Division, err error) {
	if err := h.DB.Model(&Models.Division{}).Count(&totalData).Error; err != nil {
		return 0, nil, err
	}

	// ambil fungsi offset limit Paginate di uttils
	if err := h.DB.Scopes(Utils.Paginate(page, perPage)).Find(&divisions).Error; err != nil {
		return 0, nil, err
	}

	return totalData, divisions, nil
}
