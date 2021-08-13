package handlers

import (
	"github.com/ecommerce-service/transaction-service/domain/handlers"
	"github.com/ecommerce-service/transaction-service/domain/requests"
	"github.com/ecommerce-service/transaction-service/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/pkg/response"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	HandlerContract
}

func NewTransactionHandler(handler HandlerContract) handlers.ITransactionHandler {
	return TransactionHandler{HandlerContract: handler}
}

func (h TransactionHandler) GetListForAdminWithPagination(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	transactionType := ctx.Query("transaction_type")

	uc := usecases.NewTransactionUseCase(h.UseCaseContract)
	res, pagination, err := uc.GetListForAdminWithPagination(search, orderBy, sort, transactionType, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h TransactionHandler) GetListForNormalUserWithPagination(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	status := ctx.Query("status")

	uc := usecases.NewTransactionUseCase(h.UseCaseContract)
	res, pagination, err := uc.GetListForNormalUserWithPagination(search, orderBy, sort, status, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h TransactionHandler) GetByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewTransactionUseCase(h.UseCaseContract)
	res, err := uc.GetByID(id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h TransactionHandler) Add(ctx *fiber.Ctx) (err error) {
	err = h.UseCaseContract.Config.DB.Begin()
	if err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	uc := usecases.NewTransactionUseCase(h.UseCaseContract)
	res, err := uc.Add()
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UseCaseContract.Config.DB.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h TransactionHandler) CancelPayment(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewTransactionUseCase(h.UseCaseContract)
	res, err := uc.CancelPayment(id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h TransactionHandler) ConfirmPayment(ctx *fiber.Ctx) (err error) {
	req := new(requests.ConfirmPaymentRequest)
	id := ctx.Params("id")

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	err = h.UseCaseContract.Config.DB.Begin()
	if err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	uc := usecases.NewTransactionUseCase(h.UseCaseContract)
	res, err := uc.ConfirmPayment(id)
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UseCaseContract.Config.DB.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}
