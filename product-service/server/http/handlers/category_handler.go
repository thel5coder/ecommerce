package handlers

import (
	"github.com/ecommerce-service/product-service/domain/handlers"
	"github.com/ecommerce-service/product-service/domain/requests"
	"github.com/ecommerce-service/product-service/domain/usecase"
	"github.com/ecommerce-service/product-service/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/pkg/response"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	HandlerContract
	categoryUc usecase.ICategoryUseCase
}

func NewCategoryHandler(handlerContract HandlerContract) handlers.ICategoryHandler {
	categoryUc := usecases.NewCategoryUseCase(handlerContract.UseCaseContract)

	return &CategoryHandler{
		HandlerContract: handlerContract,
		categoryUc:      categoryUc,
	}
}

func (h CategoryHandler) GetListWithPagination(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	res, pagination, err := h.categoryUc.GetListWithPagination(search, orderBy, sort, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h CategoryHandler) GetAll(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")

	res, err := h.categoryUc.GetAll(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CategoryHandler) GetByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	res, err := h.categoryUc.GetByID(id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CategoryHandler) Edit(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	req := new(requests.CategoryRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	res, err := h.categoryUc.Edit(req, id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CategoryHandler) Add(ctx *fiber.Ctx) (err error) {
	req := new(requests.CategoryRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	res, err := h.categoryUc.Add(req)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusCreated)).Send(ctx)
}

func (h CategoryHandler) Delete(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = h.categoryUc.Delete(id)

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err, http.StatusOK)).Send(ctx)
}
