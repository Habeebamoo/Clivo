package services

import "github.com/Habeebamoo/Clivo/server/internal/repositories"

type UserService interface {
	SignIn()
}

type UserSvc struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserSvc{repo}
}

func (uSvc *UserSvc) SignIn() {

}