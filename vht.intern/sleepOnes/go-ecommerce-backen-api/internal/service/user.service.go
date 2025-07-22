package service

import "go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepository(),
	}
}

func (us *UserService) GetInfoUser(id int, name string) string {
	return us.userRepo.GetInfoUser(id, name)
}
