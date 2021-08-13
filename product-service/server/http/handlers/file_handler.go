package handlers

import (
	"github.com/ecommerce-service/product-service/domain/handlers"
	"github.com/ecommerce-service/product-service/domain/usecase"
	"github.com/ecommerce-service/product-service/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/pkg/response"
	"net/http"
)

type FileHandler struct{
	HandlerContract
	fileUc usecase.IFileUseCase
}

func NewFileHandler(handlerContract HandlerContract) handlers.IFileHandler{
	fileUc := usecases.NewFileUseCase(handlerContract.UseCaseContract)

	return &FileHandler{
		HandlerContract: handlerContract,
		fileUc:          fileUc,
	}
}

func (h FileHandler) Upload(ctx *fiber.Ctx) (err error) {
	file,err := ctx.FormFile("file")
	if err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}

	res,err := h.fileUc.Upload(file)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err,http.StatusCreated)).Send(ctx)
}

func (h FileHandler) GetUrlByKey(ctx *fiber.Ctx) (err error) {
	key := ctx.Params("key")

	res,err := h.fileUc.GetUrlByKey(key)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err,http.StatusOK)).Send(ctx)
}

