package usecase

import (
	"ca_ddd_solid_dip_di/internal/application/contract/cache"
	domain_repository "ca_ddd_solid_dip_di/internal/domain/repository"
)

type GetListUserUseCase struct {
	userCache cache.Cache
	userRepo  domain_repository.IUserRepository
}

func (c *GetListUserUseCase) Execute() {

	return
}

func NewGetListUserUseCase(userRepo domain_repository.IUserRepository, userCache cache.Cache) *GetListUserUseCase {
	return &GetListUserUseCase{
		userCache: userCache,
		userRepo:  userRepo,
	}
}
