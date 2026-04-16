package usecase

import domain_repository "ca_ddd_solid_dip_di/internal/domain/repository"

type CreateUser struct {
	userRepo domain_repository.IUserRepository
}

func (c *CreateUser) Execute() {
	c.userRepo.Create()
}

func NewCreateUser(userRepo domain_repository.IUserRepository) *CreateUser {
	return &CreateUser{
		userRepo: userRepo,
	}
}
