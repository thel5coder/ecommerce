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

type ProductHandler struct {
	HandlerContract
	productUc usecase.IProductUseCase
}

//NewProductHandler function to initialize new product handler
func NewProductHandler(handlerContract HandlerContract) handlers.IProductHandler {
	productUc := usecases.NewProductUseCase(handlerContract.UseCaseContract)

	return &ProductHandler{
		HandlerContract: handlerContract,
		productUc:       productUc,
	}
}

//GetListWithPagination handler function to get list products with pagination
func (h ProductHandler) GetListWithPagination(ctx *fiber.Ctx) (err error) {
	//get query params from request
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	category := ctx.Query("category")

	//call GetListWithPagination use case function
	res, pagination, err := h.productUc.GetListWithPagination(search, orderBy, sort, category, page, limit)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

//GetByID handler function to get product by id
func (h ProductHandler) GetByID(ctx *fiber.Ctx) (err error) {
	//get params from request
	id := ctx.Params("id")

	//call GetByID use case function
	res, err := h.productUc.GetByID(id)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

//Edit handler function for edit product
func (h ProductHandler) Edit(ctx *fiber.Ctx) (err error) {
	//get params from request
	id := ctx.Params("id")

	//init new request struct
	req := new(requests.ProductRequest)

	//parse and validate the request body
	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	//initialization database transaction
	err = h.UseCaseContract.Config.DB.Begin()
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}

	//call Edit use case function
	res, err := h.productUc.Edit(req, id)
	if err != nil {
		//Rollback the database transaction when error not nil
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}

	//Commit the database transaction
	h.UseCaseContract.Config.DB.Commit()

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

//Add handler function to add new product
func (h ProductHandler) Add(ctx *fiber.Ctx) (err error) {
	//init new request struct
	req := new(requests.ProductRequest)

	//parse and validate the request body
	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	//initialization database transaction
	err = h.UseCaseContract.Config.DB.Begin()
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	//call Add use case function
	res, err := h.productUc.Add(req)
	if err != nil {
		//Rollback the database transaction when error not nil
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}
	//Commit the database transaction
	h.UseCaseContract.Config.DB.Commit()

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

//Delete handler function to delete product data
func (h ProductHandler) Delete(ctx *fiber.Ctx) (err error) {
	//get params from request
	id := ctx.Params("id")

	//initialization database transaction
	err = h.UseCaseContract.Config.DB.Begin()
	if err != nil {
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}

	//call delete use case function
	err = h.productUc.Delete(id)
	if err != nil {
		//Rollback the database transaction when error not nil
		h.UseCaseContract.Config.DB.RollBack()
		return response.NewResponse(response.NewResponseUnprocessableEntity(err)).Send(ctx)
	}

	//Commit the database transaction
	h.UseCaseContract.Config.DB.Commit()

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(nil, err, http.StatusOK)).Send(ctx)
}
