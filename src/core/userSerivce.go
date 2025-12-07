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

func (us *userService) GetUserByLogin(login string) (*shared.User, error) {
	if len(login) < 3 {
		return nil, UserDoesntExits
	}
	u, err := us.repository.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}
	if u.DeletedAt != nil {
		return nil, UserDoesntExits
	}
	return u, nil
}

//func (us *userService) GetUserByEmail(email string) (*shared.User, error) {
//
//}

//func (us *userService) CreateUser(dto *shared.UserInputDto) (*shared.User, error) {
//	_, err := us.GetUserByLogin(dto.Login)
//	if err != nil && err != UserDoesntExits {
//		nil, err
//	}
//	var u = shared.User{
//		Login: dto.Login,
//		e
//	}
//	return &user, nil
//}
