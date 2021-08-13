package routers

import (
	"github.com/ecommerce-service/product-service/server/http/handlers"
	"github.com/ecommerce-service/product-service/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ProductRouters struct{
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewProductRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &ProductRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r ProductRouters) RegisterRouter() {
	handler := handlers.NewProductHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	productRouters := r.RouteGroup.Group("/product")
	productRouters.Use(jwt.Use)
	productRouters.Get("", handler.GetListWithPagination)
	productRouters.Get("/:id", handler.GetByID)
	productRouters.Put("/:id", handler.Edit)
	productRouters.Post("", handler.Add)
	productRouters.Delete("/:id", handler.Delete)
}

