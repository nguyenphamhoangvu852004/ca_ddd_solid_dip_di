package domain_repository

type IUserRepository interface {
	Get(userID string) (interface{}, error)
	Create()
}
