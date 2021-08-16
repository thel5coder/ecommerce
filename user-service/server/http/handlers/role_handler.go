package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/ecommerce/user-service/domain/handlers"
	"github.com/thel5coder/ecommerce/user-service/usecases"
	"github.com/thel5coder/pkg/response"
	"net/http"
)

type RoleHandler struct{
	Handler HandlerContract
}

func NewRoleHandler(handler HandlerContract) handlers.IRoleHandler{
	return &RoleHandler{Handler: handler}
}

func (h RoleHandler) BrowseAll(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")

	uc := usecases.NewRoleUseCase(h.Handler.UseCaseContract)
	res,err := uc.BrowseAll(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err,http.StatusOK)).Send(ctx)
}

