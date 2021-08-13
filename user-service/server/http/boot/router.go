package boot

import (
	"github.com/ecommerce-service/user-service/server/http/boot/routers"
	"github.com/ecommerce-service/user-service/server/http/handlers"
	"github.com/ecommerce-service/user-service/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
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

	//register role routers
	roleRouters := routers.NewRoleRouters(rootRouter, handler)
	roleRouters.RegisterRouter()

	//register user routers
	userRouters := routers.NewUserRouters(rootRouter, handler)
	userRouters.RegisterRouter()

	//register authentication routers
	authenticationRouters := routers.NewAuthenticationRouters(rootRouter, handler)
	authenticationRouters.RegisterRouter()
}
