package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/ecommerce/product-service/server/http/handlers"
	"github.com/thel5coder/ecommerce/product-service/server/http/middlewares"
)

type ProductRouters struct{
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

//NewProductRouters function to initialize new routers
func NewProductRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &ProductRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

//RegisterRouter register category routers
func (r ProductRouters) RegisterRouter() {
	//initialize the category handler
	handler := handlers.NewProductHandler(r.Handler)

	//initialize jwt middleware
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	//creating product routing group
	productRouters := r.RouteGroup.Group("/product")

	//secure the route with jwt middleware
	productRouters.Use(jwt.Use)

	//routing
	productRouters.Get("", handler.GetListWithPagination)
	productRouters.Get("/:id", handler.GetByID)
	productRouters.Put("/:id", handler.Edit)
	productRouters.Post("", handler.Add)
	productRouters.Delete("/:id", handler.Delete)
}

