package usecase

import (
	"ca_ddd_solid_dip_di/internal/application/contract/cache"
	domain_repository "ca_ddd_solid_dip_di/internal/domain/repository"
)

type CreateUserUseCase struct {
	userCache cache.Cache
	userRepo  domain_repository.IUserRepository
}

func (c *CreateUserUseCase) Execute() {

	// call DB to get data
	data, err := c.userRepo.Get("1")
	if err != nil {
		panic(err)
	}

	if data != nil {
		// return data from cache
		return
	}

	c.userRepo.Create()

	return
}

func NewCreateUserUseCase(userCache cache.Cache, userRepo domain_repository.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userCache: userCache,
		userRepo:  userRepo,
	}
}
