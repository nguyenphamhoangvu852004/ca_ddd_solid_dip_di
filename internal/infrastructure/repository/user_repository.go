package infrastructure_repository

import domain_repository "ca_ddd_solid_dip_di/internal/domain/repository"

type UserRepositoryImpl struct {
	// Mysql
}

// Create implements [domain_repository.IUserRepository].
func (u *UserRepositoryImpl) Create() {
	panic("unimplemented")
}

// Get implements [domain_repository.IUserRepository].
func (u *UserRepositoryImpl) Get() {
	panic("unimplemented")
}

func NewUserRepostoryImpl() domain_repository.IUserRepository {
	return &UserRepositoryImpl{}
}
