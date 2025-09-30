package core

import "blog-api/src/repository"

func Init() bool {
	UserService = newUserService(repository.Repo)
	return true
}
