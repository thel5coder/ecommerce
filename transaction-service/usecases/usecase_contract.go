package usecases

import (
	"github.com/ecommerce-service/transaction-service/domain/configs"
	"github.com/ecommerce-service/transaction-service/domain/view_models"
)

type UseCaseContract struct {
	RequestID string
	UserID    string
	RoleID    int
	Config    *configs.Config
}

func NewUseCaseContract(requestID string, config *configs.Config) *UseCaseContract {
	return &UseCaseContract{RequestID: requestID, Config: config}
}

const (
	//default limit for pagination
	defaultLimit = 10

	//max limit for pagination
	maxLimit = 50

	//default order by
	defaultOrderBy = "id"

	//default sort
	defaultSort = "asc"

	//default last page for pagination
	defaultLastPage = 0

	// DefaultIDNormalUsers default role id for normal users
	DefaultIDNormalUsers = 2

	// DefaultTransactionType default transaction type
	DefaultTransactionType = "on_going"

	// SuccessTransactionType transaction success type
	SuccessTransactionType = "success"

	// CancelTransactionType transaction canceled type
	CancelTransactionType = "canceled"
)

func (uc *UseCaseContract) SetPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc *UseCaseContract) SetPaginationResponse(page, limit, total int) (res view_models.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	vm := view_models.NewPaginationVm()
	res = vm.Build(view_models.DetailPaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	})

	return res
}
