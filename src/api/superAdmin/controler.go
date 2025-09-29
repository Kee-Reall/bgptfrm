package superAdmin

import (
	"blog-api/src/repository"
	"blog-api/src/shared"
	"blog-api/src/utils"
	"encoding/json"
	"net/http"
)

type UserProvider interface {
	GetUsers(params shared.PaginationParams) ([]shared.User, error)
}

func sortByExtractor(_ []string) string {
	return "createdAt"
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.Repo.GetUsers(utils.ParsePaginationParams(r.URL.Query(), sortByExtractor))
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
