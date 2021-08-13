package usecases

import (
	"github.com/ecommerce-service/transaction-service/domain/requests"
	"github.com/ecommerce-service/transaction-service/domain/view_models"
)

type ICartUseCase interface {
	GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CartVm, pagination view_models.PaginationVm, err error)

	GetAllByUserId(userId string) (res []view_models.CartVm, err error)

	GetByID(id string) (res view_models.CartVm, err error)

	GetBy(column, operator string, value interface{}) (res view_models.CartVm, err error)

	Edit(req *requests.CartRequest, id string) (res string, err error)

	Add(req *requests.CartRequest) (res string, err error)

	Delete(id string) (err error)

	DeleteAllByUserId() (err error)

	Count(search string) (res int, err error)

	CountBy(column, operator string, value interface{}) (res int, err error)
}
