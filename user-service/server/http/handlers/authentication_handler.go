package handlers

import (
	"github.com/ecommerce/user-service/domain/handlers"
	"github.com/ecommerce/user-service/domain/requests"
	"github.com/ecommerce/user-service/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/pkg/response"
	"net/http"
)

type AuthenticationHandler struct {
	HandlerContract
}

func NewAuthenticationHandler(handlerContract HandlerContract) handlers.IAuthenticationHandler {
	return &AuthenticationHandler{HandlerContract: handlerContract}
}

func (h AuthenticationHandler) Login(ctx *fiber.Ctx) error {
	req := new(requests.LoginRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewAuthenticationUseCase(h.UseCaseContract)
	res, err := uc.Login(req)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h AuthenticationHandler) Register(ctx *fiber.Ctx) error {
	req := new(requests.RegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewAuthenticationUseCase(h.UseCaseContract)
	err := uc.Registration(req)

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err, http.StatusOK)).Send(ctx)
}
