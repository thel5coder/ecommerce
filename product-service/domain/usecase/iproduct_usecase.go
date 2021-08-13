package usecase

import (
	"github.com/ecommerce/product-service/domain/requests"
	"github.com/ecommerce/product-service/domain/view_models"
)

type IProductUseCase interface {
	GetListWithPagination(search, orderBy, sort,category string, page, limit int) (res []view_models.ListProductVm, pagination view_models.PaginationVm, err error)

	GetByID(id string) (res view_models.DetailProductVm, err error)

	Edit(req *requests.ProductRequest, id string) (res string, err error)

	Add(req *requests.ProductRequest) (res string, err error)

	Delete(id string) (err error)

	Count(search,category string) (res int, err error)

	CountBy(column, operator, id string, value interface{}) (res int, err error)
}
