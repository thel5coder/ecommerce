package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/ecommerce/transaction-service/server/http/handlers"
	"github.com/thel5coder/ecommerce/transaction-service/server/http/middlewares"
)

type CartRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewCartRouters(rootGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &CartRouters{
		RouteGroup: rootGroup,
		Handler:    handler,
	}
}

func (r CartRouters) RegisterRouter() {
	handler := handlers.NewCartHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)
	normalUserMiddleware := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	cartRouters := r.RouteGroup.Group("/cart")
	cartRouters.Use(jwt.Use)
	cartRouters.Use(normalUserMiddleware.Use)
	cartRouters.Get("", handler.GetListWithPagination)
	cartRouters.Get("/:id", handler.GetByID)
	cartRouters.Put("/:id", handler.Edit)
	cartRouters.Post("", handler.Add)
	cartRouters.Delete("/:id", handler.Delete)
}
