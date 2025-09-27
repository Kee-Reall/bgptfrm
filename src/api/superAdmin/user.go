package superAdmin

import (
	"blog-api/src/api/common"
)

type SuperAdminUserSearchParams struct {
	common.PaginationParams
	SortingKey      string `json:"sort_by"`
	SearchLoginTerm string `json:"searchLoginTerm"`
	SearchEmailTerm string `json:"searchEmailTerm"`
}

type SuperAdminUserController struct {
}

var controller *SuperAdminUserController = nil

func GetController() *SuperAdminUserController {
	if controller == nil {
		controller = &SuperAdminUserController{}

	}
	return controller
}

func (c *SuperAdminUserController) getUsers() {

}
