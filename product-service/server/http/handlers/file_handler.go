package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/ecommerce/product-service/domain/handlers"
	"github.com/thel5coder/ecommerce/product-service/domain/usecase"
	"github.com/thel5coder/ecommerce/product-service/usecases"
	"github.com/thel5coder/pkg/response"
	"net/http"
)

type FileHandler struct{
	HandlerContract
	fileUc usecase.IFileUseCase
}

//NewFileHandler function to initialize new file handler
func NewFileHandler(handlerContract HandlerContract) handlers.IFileHandler{
	fileUc := usecases.NewFileUseCase(handlerContract.UseCaseContract)

	return &FileHandler{
		HandlerContract: handlerContract,
		fileUc:          fileUc,
	}
}

//Upload handler function to handling upload request
func (h FileHandler) Upload(ctx *fiber.Ctx) (err error) {
	//get file from request and check it
	file,err := ctx.FormFile("file")
	if err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}

	//call Upload use case function
	res,err := h.fileUc.Upload(file)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res,err,http.StatusCreated)).Send(ctx)
}

//GetUrlByKey handler function to handling get file from min.io
func (h FileHandler) GetUrlByKey(ctx *fiber.Ctx) (err error) {
	//get params from request
	key := ctx.Params("key")

	//call Upload use case function
	res,err := h.fileUc.GetUrlByKey(key)

	//return response with response factory
	return response.NewResponse(response.NewResponseWithOutMeta(res,err,http.StatusOK)).Send(ctx)
}

