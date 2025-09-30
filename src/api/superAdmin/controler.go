package superAdmin

import (
	"blog-api/src/core"
	"blog-api/src/shared"
	"blog-api/src/utils"
	"blog-api/src/utils/validation"
	"encoding/json"
	"net/http"
)

type UserProvider interface {
	GetUsers(params shared.PaginationParams) ([]shared.User, error)
}

func sortByExtractor(input []string) (output string) {
	output = "createdAt"

	if len(input) <= 0 || input[0] == "" {
		return
	}

	availableKey := [3]string{"id", "login", "email"}
	for _, key := range availableKey {
		if input[0] == key {
			output = key
			return
		}
	}

	return
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	params := utils.ParsePaginationParams(r.URL.Query(), sortByExtractor)
	users, err := core.UserService.GetUsers(&params)
	if err != nil {
		utils.WriteException(w)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		utils.WriteException(w)
		return
	}
	w.WriteHeader(200)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var dto shared.UserInputDto
	if errs := validation.BindAndValidate(r.Body, &dto); errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errs)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}
