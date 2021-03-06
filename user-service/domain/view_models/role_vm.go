package view_models

import "github.com/thel5coder/ecommerce/user-service/domain/models"

type RoleVm struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewRoleVm(model *models.Role) RoleVm {
	return RoleVm{
		ID:   model.Id(),
		Name: model.Name(),
	}
}
