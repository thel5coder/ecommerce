package usecases

import (
	"github.com/thel5coder/ecommerce/transaction-service/domain/requests"
)

type ITransactionDetailUseCase interface {

	Add(req requests.TransactionDetailRequest,transactionId string) (err error)

	Store(reqs []requests.TransactionDetailRequest,transactionId string) (err error)
}
