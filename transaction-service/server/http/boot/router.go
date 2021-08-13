package boot

import (
	"github.com/ecommerce-service/transaction-service/server/http/boot/routers"
	"github.com/ecommerce-service/transaction-service/server/http/handlers"
	"github.com/ecommerce-service/transaction-service/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

func (boot *Boot) RegisterAllRouters() {
	handler := handlers.HandlerContract{
		UseCaseContract: usecases.NewUseCaseContract(xid.New().String(), boot.Config),
		App: boot.App,
	}

	//init route group
	rootRouter := boot.App.Group("/api")
	//check health
	rootRouter.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("it's working")
	})

	//cart routers
	cartRouters := routers.NewCartRouters(rootRouter, handler)
	cartRouters.RegisterRouter()

	// Transaction routers
	transactionRouters := routers.NewTransactionRouters(rootRouter, handler)
	transactionRouters.RegisterRouter()
}
