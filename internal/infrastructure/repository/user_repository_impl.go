package infrastructure_repository

import domain_repository "ca_ddd_solid_dip_di/internal/domain/repository"

type UserRepositoryImpl struct {
	// Mysql
}

// Get implements [domain_repository.IUserRepository].
func (u *UserRepositoryImpl) Get(userID string) (interface{}, error) {
	panic("unimplemented")
}

// Create implements [domain_repository.IUserRepository].
func (u *UserRepositoryImpl) Create() {
	panic("unimplemented")
}

func NewUserRepostoryImpl() domain_repository.IUserRepository {
	return &UserRepositoryImpl{}
}
