package usecases

import (
	"github.com/ecommerce/transaction-service/domain/requests"
)

type ITransactionDetailUseCase interface {

	Add(req requests.TransactionDetailRequest,transactionId string) (err error)

	Store(reqs []requests.TransactionDetailRequest,transactionId string) (err error)
}
