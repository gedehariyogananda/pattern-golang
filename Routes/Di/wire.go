//go:build wireinject
// +build wireinject

package Di

import (
	"github.com/gedehariyogananda/pattern-golang/Controllers"
	"github.com/gedehariyogananda/pattern-golang/Repositories"
	"github.com/gedehariyogananda/pattern-golang/Services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func DIAuth(db *gorm.DB) *Controllers.AuthController {
	panic(wire.Build(wire.NewSet(
		Repositories.AuthRepositoryProvider,
		Services.AuthServiceProvider,
		Controllers.AuthControllerProvider,

		wire.Bind(new(Controllers.IAuthController), new(*Controllers.AuthController)),
		wire.Bind(new(Services.IAuthService), new(*Services.AuthService)),
		wire.Bind(new(Repositories.IAuthRepository), new(*Repositories.AuthRepository)),
	),
	))

	return &Controllers.AuthController{}
}
