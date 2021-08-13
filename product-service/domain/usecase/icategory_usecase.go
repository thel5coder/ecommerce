package usecase

import (
	"github.com/ecommerce/product-service/domain/requests"
	"github.com/ecommerce/product-service/domain/view_models"
)

type ICategoryUseCase interface {
	GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CategoryVm, pagination view_models.PaginationVm, err error)

	GetAll(search string) (res []view_models.CategoryVm, err error)

	GetByID(id string) (res view_models.CategoryVm, err error)

	Edit(req *requests.CategoryRequest, id string) (res string, err error)

	Add(req *requests.CategoryRequest) (res string, err error)

	Delete(id string) (err error)

	Count(search string) (res int, err error)

	CountBy(column, operator, id string, value interface{}) (res int, err error)
}
