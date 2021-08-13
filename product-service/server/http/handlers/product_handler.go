package handlers

import (
	"github.com/ecommerce/product-service/domain/handlers"
	"github.com/ecommerce/product-service/domain/requests"
	"github.com/ecommerce/product-service/domain/usecase"
	"github.com/ecommerce/product-service/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/pkg/response"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	HandlerContract
	productUc usecase.IProductUseCase
}

func NewProductHandler(handlerContract HandlerContract) handlers.IProductHandler {
	productUc := usecases.NewProductUseCase(handlerContract.UseCaseContract)

	return &ProductHandler{
		HandlerContract: handlerContract,
		productUc:       productUc,
	}
}

func (h ProductHandler) GetListWithPagination(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	category := ctx.Query("category")

	res, pagination, err := h.productUc.GetListWithPagination(search, orderBy, sort, category, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h ProductHandler) GetByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	res, err := h.productUc.GetByID(id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h ProductHandler) Edit(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	req := new(requests.ProductRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	err = h.UseCaseContract.Config.DB.Begin()
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	res, err := h.productUc.Edit(req, id)
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UseCaseContract.Config.DB.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h ProductHandler) Add(ctx *fiber.Ctx) (err error) {
	req := new(requests.ProductRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	err = h.UseCaseContract.Config.DB.Begin()
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	res, err := h.productUc.Add(req)
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UseCaseContract.Config.DB.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h ProductHandler) Delete(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = h.UseCaseContract.Config.DB.Begin()
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	err = h.productUc.Delete(id)
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	h.UseCaseContract.Config.DB.Commit()

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err, http.StatusOK)).Send(ctx)
}
