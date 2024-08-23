package Services

import (
	"github.com/gedehariyogananda/pattern-golang/Models/Dto"
	"github.com/gedehariyogananda/pattern-golang/Repositories"
)

type (
	IAuthService interface {
		Register(request *Dto.RegisterRequest) (err error)
	}

	AuthService struct {
		repo Repositories.IAuthRepository
	}
)

func AuthServiceProvider(repo Repositories.IAuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (h *AuthService) Register(request *Dto.RegisterRequest) (err error) {
	if err := h.repo.CheckUniqueField(request); err != nil {
		return err
	}

	if err := h.repo.InsertForRegister(request); err != nil {
		return err
	}

	return nil

}
