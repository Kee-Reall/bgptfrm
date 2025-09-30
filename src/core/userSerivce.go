package core

import (
	"blog-api/src/repository"
	"blog-api/src/shared"
)

type userService struct {
	repository *repository.Repository
}

var UserService *userService = nil

func newUserService(repo *repository.Repository) *userService {
	return &userService{repo}
}

func (us *userService) GetUsers(params *shared.PaginationParams) ([]shared.User, error) {
	return us.repository.GetUsers(params)
}

func (us *userService) CreateUser(dto *shared.UserInputDto) {

}
