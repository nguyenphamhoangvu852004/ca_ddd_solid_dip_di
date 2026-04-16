package dependency_injection

import (
	"ca_ddd_solid_dip_di/internal/controller"
	infrastructure_repository "ca_ddd_solid_dip_di/internal/infrastructure/repository"
	"ca_ddd_solid_dip_di/internal/usecase"
)

// UserDI provides all dependencies for User module
type UserDI struct {
	UserController *controller.UserController
}

// InitUserDependencies initializes all User module dependencies
func InitUserDependencies() *UserDI {
	// Initialize repository (infrastructure layer)
	userRepository := infrastructure_repository.NewUserRepostoryImpl()

	// Initialize use case (application layer)
	createUserUseCase := usecase.NewCreateUser(userRepository)

	// Initialize controller (presentation layer)
	userController := controller.NewUserController(createUserUseCase)

	return &UserDI{
		UserController: userController,
	}
}
