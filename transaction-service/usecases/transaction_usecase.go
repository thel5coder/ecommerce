package usecases

import (
	"fmt"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/ecommerce-service/transaction-service/domain/requests"
	"github.com/ecommerce-service/transaction-service/domain/usecases"
	"github.com/ecommerce-service/transaction-service/domain/view_models"
	"github.com/ecommerce-service/transaction-service/repository/commands"
	"github.com/ecommerce-service/transaction-service/repository/queries"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
	"time"
)

type TransactionUseCase struct {
	*UseCaseContract
}

func NewTransactionUseCase(useCaseContract *UseCaseContract) usecases.ITransactionUseCase {
	return &TransactionUseCase{UseCaseContract: useCaseContract}
}

func (uc TransactionUseCase) GetListForAdminWithPagination(search, orderBy, sort, status string, page, limit int) (res []view_models.TransactionListVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	transactions, err := q.Browse(search, orderBy, sort, status, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-browse")
		return res, pagination, err
	}
	for _, transaction := range transactions.([]*models.Transactions) {
		res = append(res, view_models.NewTransactionListVm(transaction))
	}

	//set pagination
	totalCount, err := uc.Count(search, "", status)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transaction-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc TransactionUseCase) GetListForNormalUserWithPagination(search, orderBy, sort, status string, page, limit int) (res []view_models.TransactionListVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	transactions, err := q.BrowseByUserId(search, orderBy, sort, uc.UserID, status, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-browse")
		return res, pagination, err
	}
	for _, transaction := range transactions.([]*models.Transactions) {
		res = append(res, view_models.NewTransactionListVm(transaction))
	}

	//set pagination
	totalCount, err := uc.Count(search, uc.UserID, status)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transaction-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc TransactionUseCase) GetByID(id string) (res view_models.TransactionDetailVm, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)

	transaction, err := q.ReadBy("t.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-readByID")
		return res, err
	}
	res = view_models.NewTransactionDetailVm(transaction.(*models.Transactions))

	return res, nil
}

func (uc TransactionUseCase) CancelPayment(id string) (res string, err error) {
	now := time.Now().UTC()

	model := models.NewTransactionModel().SetId(id).SetStatus(CancelTransactionType).SetCanceledAt(now).SetUpdatedAt(now)
	cmd := commands.NewTransactionCommand(uc.Config.DB, model)
	res, err = cmd.EditCancelPayment()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-editCancelPayment")
		return res, err
	}

	return res, nil
}

func (uc TransactionUseCase) ConfirmPayment(id string) (res string, err error) {
	now := time.Now().UTC()

	model := models.NewTransactionModel().SetId(id).SetStatus(SuccessTransactionType).SetPaidAt(now).SetUpdatedAt(now)
	cmd := commands.NewTransactionCommand(uc.Config.DB, model)
	err = cmd.EditPaymentReceived()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-editPaymentReceived")
		return res, err
	}
	res = id

	return res, nil
}

func (uc TransactionUseCase) Add() (res string, err error) {
	now := time.Now().UTC()

	cartUc := NewCartUseCase(uc.UseCaseContract)
	carts, err := cartUc.GetAllByUserId(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-getAllByUserId")
		return res, err
	}

	totalAmount, transactionDetailRequest := uc.GetTotalAmountAndBuildTransactionDetailRequest(carts)
	transactionNumber := uc.GetTransactionNumber()
	model := models.NewTransactionModel().SetStatus(DefaultTransactionType).SetTransactionNumber(transactionNumber).
		SetTotal(totalAmount).SetCreatedAt(now).SetUpdatedAt(now).SetUserId(uc.UserID)
	cmd := commands.NewTransactionCommand(uc.Config.DB, model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-add")
		return res, err
	}

	transactionDetailUc := NewTransactionDetailUseCase(uc.UseCaseContract)
	err = transactionDetailUc.Store(transactionDetailRequest, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transactionDetail-store")
		return res, err
	}

	err = cartUc.DeleteAllByUserId()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-deleteAllByUserId")
		return res, err
	}

	return res, nil
}

func (uc TransactionUseCase) Count(search, userId, status string) (res int, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)

	res, err = q.Count(search, userId, status)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-count")
		return res, err
	}

	return res, nil
}

func (uc TransactionUseCase) CountAll() (res int, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)

	res, err = q.CountAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-countAll")
		return res, err
	}

	return res, nil
}

func (uc TransactionUseCase) GetTransactionNumber() (res string) {
	count, err := uc.CountAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transaction-countAll")
		return res
	}
	res = time.Now().UTC().Format("2006-01-02")
	res += fmt.Sprintf("%04d", count+1)

	return res
}

func (uc TransactionUseCase) GetTotalAmountAndBuildTransactionDetailRequest(carts []view_models.CartVm) (totalAmount float64, req []requests.TransactionDetailRequest) {
	for _, cart := range carts {
		totalAmount += cart.SubTotal
		req = append(req, requests.TransactionDetailRequest{
			Name:     cart.Name,
			Sku:      cart.Sku,
			Category: cart.Category,
			Price:    cart.Price,
			Discount: 0,
			Quantity: int(cart.Quantity),
			SubTotal: cart.SubTotal,
		})
	}

	return totalAmount, req
}
