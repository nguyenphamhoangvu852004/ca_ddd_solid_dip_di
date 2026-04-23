package initialize

import (
	auth_usecase "ca_ddd_solid_dip_di/internal/application/usecase/auth"
	user_usecase "ca_ddd_solid_dip_di/internal/application/usecase/user"
	"ca_ddd_solid_dip_di/internal/infrastructure/cache"
	infrastructure_repository "ca_ddd_solid_dip_di/internal/infrastructure/repository"
	controller "ca_ddd_solid_dip_di/internal/presentation"
)

type AppManagement struct {
	UserController *controller.UserController
	AuthController *controller.AuthController
}

func InitDependencies() *AppManagement {
	//Implementation in infrastructure
	userRepositoryImpl := infrastructure_repository.NewUserRepostoryImpl()
	RedisCacheImpl := cache.NewRedisCacheImpl()

	//usecase
	createUserUseCase := user_usecase.NewCreateUserUseCase(RedisCacheImpl, userRepositoryImpl)
	getUsersUseCase := user_usecase.NewGetListUserUseCase(userRepositoryImpl, RedisCacheImpl)

	registrateUseCase := auth_usecase.NewRegistrateUseCase(userRepositoryImpl)

	//controller
	userController := controller.NewUserController(createUserUseCase, getUsersUseCase)
	authController := controller.NewAuthController(registrateUseCase)

	return &AppManagement{
		UserController: userController,
		AuthController: authController,
	}
}
