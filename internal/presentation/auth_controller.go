package presentation

import (
	usecase "ca_ddd_solid_dip_di/internal/application/usecase/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	registrateUseCase *usecase.RegistrateUseCase
}

func (ac *AuthController) Registrate(ctx *gin.Context) {
	ac.registrateUseCase.Execute()
	return
}

func NewAuthController(registrateUseCase *usecase.RegistrateUseCase) *AuthController {
	return &AuthController{
		registrateUseCase: registrateUseCase,
	}
}
