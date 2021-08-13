package usecase

import "github.com/ecommerce/user-service/domain/view_models"

type IRoleUseCase interface {
	BrowseAll(search string) (res []view_models.RoleVm, err error)
}
