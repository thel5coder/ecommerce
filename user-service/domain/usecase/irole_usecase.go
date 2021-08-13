package usecase

import "github.com/ecommerce-service/user-service/domain/view_models"

type IRoleUseCase interface {
	BrowseAll(search string) (res []view_models.RoleVm, err error)
}
