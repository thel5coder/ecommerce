package routers

import (
	"github.com/ecommerce-service/user-service/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewAuthenticationRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &AuthenticationRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r AuthenticationRouters) RegisterRouter() {
	handler := handlers.NewAuthenticationHandler(r.Handler)

	authenticationRouters := r.RouteGroup.Group("/auth")
	authenticationRouters.Post("/login", handler.Login)
	authenticationRouters.Post("/register", handler.Register)
}
