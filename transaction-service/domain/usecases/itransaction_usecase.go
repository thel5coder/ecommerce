package usecases

import (
	"github.com/ecommerce-service/transaction-service/domain/requests"
	"github.com/ecommerce-service/transaction-service/domain/view_models"
)

type ITransactionUseCase interface {
	GetListForAdminWithPagination(search, orderBy, sort, status string, page, limit int) (res []view_models.TransactionListVm, pagination view_models.PaginationVm, err error)

	GetListForNormalUserWithPagination(search, orderBy, sort, status string, page, limit int) (res []view_models.TransactionListVm, pagination view_models.PaginationVm, err error)

	GetByID(id string) (res view_models.TransactionDetailVm, err error)

	CancelPayment(id string) (res string, err error)

	ConfirmPayment(id string) (res string, err error)

	Add() (res string, err error)

	Count(search, userId, status string) (res int, err error)

	CountAll() (res int, err error)

	GetTransactionNumber() string

	GetTotalAmountAndBuildTransactionDetailRequest(carts []view_models.CartVm) (totalAmount float64,req []requests.TransactionDetailRequest)
}
