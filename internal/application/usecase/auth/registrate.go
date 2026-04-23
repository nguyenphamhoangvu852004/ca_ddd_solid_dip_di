package usecase

import domain_repository "ca_ddd_solid_dip_di/internal/domain/repository"

type RegistrateUseCase struct {
	userRepo domain_repository.IUserRepository
}

func (c *RegistrateUseCase) Execute() {

	return
}

func NewRegistrateUseCase(userRepo domain_repository.IUserRepository) *RegistrateUseCase {
	return &RegistrateUseCase{
		userRepo: userRepo,
	}
}
