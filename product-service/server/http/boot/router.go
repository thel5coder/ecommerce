package boot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"github.com/thel5coder/ecommerce/product-service/server/http/boot/routers"
	"github.com/thel5coder/ecommerce/product-service/server/http/handlers"
	"github.com/thel5coder/ecommerce/product-service/usecases"
)

func (boot *Boot) RegisterAllRouters() {
	handler := handlers.HandlerContract{
		UseCaseContract: usecases.NewUseCaseContract(xid.New().String(), boot.Config),
		App:             boot.App,
	}

	//init route group
	rootRouter := boot.App.Group("/api")
	//check health
	rootRouter.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("it's working")
	})

	//register category routers
	categoryRouters := routers.NewCategoryRouters(rootRouter, handler)
	categoryRouters.RegisterRouter()

	//register file routers
	fileRouters := routers.NewFileRouters(rootRouter, handler)
	fileRouters.RegisterRouter()

	//register product routers
	productRouters := routers.NewProductRouters(rootRouter,handler)
	productRouters.RegisterRouter()
}
