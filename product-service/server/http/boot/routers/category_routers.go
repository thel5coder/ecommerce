package routers

import (
	"github.com/ecommerce/product-service/server/http/handlers"
	"github.com/ecommerce/product-service/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type CategoryRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

//NewCategoryRouters function to initialize new routers
func NewCategoryRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &CategoryRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

//RegisterRouter register category routers
func (r CategoryRouters) RegisterRouter() {
	//initialize the category handler
	handler := handlers.NewCategoryHandler(r.Handler)

	//initialize jwt middleware
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	//creating category routing group
	categoryRouters := r.RouteGroup.Group("/category")

	//secure the route with jwt middleware
	categoryRouters.Use(jwt.Use)

	//routing
	categoryRouters.Get("", handler.GetListWithPagination)
	categoryRouters.Get("/all", handler.GetAll)
	categoryRouters.Get("/:id", handler.GetByID)
	categoryRouters.Put("/:id", handler.Edit)
	categoryRouters.Post("", handler.Add)
	categoryRouters.Delete("/:id", handler.Delete)
}
