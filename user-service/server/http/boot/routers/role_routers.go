package routers

import (
	"github.com/ecommerce/user-service/server/http/handlers"
	"github.com/ecommerce/user-service/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type RoleRouters struct{
	RouteGroup fiber.Router
	Handler handlers.HandlerContract
}

func NewRoleRouters(routeGroup fiber.Router,handler handlers.HandlerContract) IRouters{
	return &RoleRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r RoleRouters) RegisterRouter() {
	handler := handlers.NewRoleHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	roleRouters := r.RouteGroup.Group("/role")
	roleRouters.Use(jwt.Use)
	roleRouters.Get("",handler.BrowseAll)
}

