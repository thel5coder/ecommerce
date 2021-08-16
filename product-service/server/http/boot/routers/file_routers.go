package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/ecommerce/product-service/server/http/handlers"
	"github.com/thel5coder/ecommerce/product-service/server/http/middlewares"
)

type FileRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

//NewFileRouters function to initialize new routers
func NewFileRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &FileRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

//RegisterRouter register category routers
func (r FileRouters) RegisterRouter() {
	//initialize the file handler
	handler := handlers.NewFileHandler(r.Handler)

	//initialize jwt middleware
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	//creating file routing group
	fileRouters := r.RouteGroup.Group("/file")

	//secure the route with jwt middleware
	fileRouters.Use(jwt.Use)

	//routing
	fileRouters.Get("/:key", handler.GetUrlByKey)
	fileRouters.Post("", handler.Upload)
}
