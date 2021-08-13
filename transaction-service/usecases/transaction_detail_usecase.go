package usecases

import (
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/ecommerce-service/transaction-service/domain/requests"
	"github.com/ecommerce-service/transaction-service/domain/usecases"
	"github.com/ecommerce-service/transaction-service/repository/commands"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
	"time"
)

type TransactionDetailUseCase struct {
	*UseCaseContract
}

func NewTransactionDetailUseCase(useCaseContract *UseCaseContract) usecases.ITransactionDetailUseCase {
	return &TransactionDetailUseCase{UseCaseContract: useCaseContract}
}

func (uc TransactionDetailUseCase) Add(req requests.TransactionDetailRequest, transactionId string) (err error) {
	now := time.Now().UTC()

	model := models.NewTransactionDetailModel().SetTransactionId(transactionId).
		SetName(req.Name).SetSku(req.Sku).SetCategory(req.Category).
		SetPrice(req.Price).SetQuantity(int64(req.Quantity)).SetSubTotal(req.SubTotal).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := commands.NewTransactionDetailCommand(uc.Config.DB, model)
	err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-transactionDetail-add")
		return err
	}

	return nil
}

func (uc TransactionDetailUseCase) Store(reqs []requests.TransactionDetailRequest, transactionId string) (err error) {
	for _, req := range reqs {
		err = uc.Add(req, transactionId)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transactionDetail-add")
			return err
		}
	}

	return nil
}
