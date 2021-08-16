package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/ecommerce/product-service/domain/handlers"
	"github.com/thel5coder/ecommerce/product-service/domain/requests"
	"github.com/thel5coder/ecommerce/product-service/domain/usecase"
	"github.com/thel5coder/ecommerce/product-service/usecases"
	"github.com/thel5coder/pkg/response"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	HandlerContract
	categoryUc usecase.ICategoryUseCase
}

//NewCategoryHandler function to initialize new category handler
func NewCategoryHandler(handlerContract HandlerContract) handlers.ICategoryHandler {
	categoryUc := usecases.NewCategoryUseCase(handlerContract.UseCaseContract)

	return &CategoryHandler{
		HandlerContract: handlerContract,
		categoryUc:      categoryUc,
	}
}

//GetListWithPagination handler function to get list category with pagination
func (h CategoryHandler) GetListWithPagination(ctx *fiber.Ctx) (err error) {
	//get query params from request
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	//call GetListWithPagination use case function
	res, pagination, err := h.categoryUc.GetListWithPagination(search, orderBy, sort, page, limit)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

//GetAll handler function to get all category data
func (h CategoryHandler) GetAll(ctx *fiber.Ctx) (err error) {
	//get query params from request
	search := ctx.Query("search")

	//call GetAll use case function
	res, err := h.categoryUc.GetAll(search)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

//GetByID handler function to get category by id
func (h CategoryHandler) GetByID(ctx *fiber.Ctx) (err error) {
	//get params from request
	id := ctx.Params("id")

	//call GetById use case function
	res, err := h.categoryUc.GetByID(id)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

//Edit handler function for edit product
func (h CategoryHandler) Edit(ctx *fiber.Ctx) (err error) {
	//get params from request
	id := ctx.Params("id")

	//init new request struct
	req := new(requests.CategoryRequest)

	//parse and validate the request body
	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	//call Edit use case function
	res, err := h.categoryUc.Edit(req, id)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

//Add handler function to add new product
func (h CategoryHandler) Add(ctx *fiber.Ctx) (err error) {
	//init new request struct
	req := new(requests.CategoryRequest)

	//parse and validate the request body
	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	//call Add use case function
	res, err := h.categoryUc.Add(req)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusCreated)).Send(ctx)
}

//Delete handler function to delete product data
func (h CategoryHandler) Delete(ctx *fiber.Ctx) (err error) {
	//get params from request
	id := ctx.Params("id")

	//call delete use case function
	err = h.categoryUc.Delete(id)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(nil, err, http.StatusOK)).Send(ctx)
}
