package routers

import (
	"github.com/ecommerce-service/transaction-service/server/http/handlers"
	"github.com/ecommerce-service/transaction-service/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type TransactionRoutes struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewTransactionRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &TransactionRoutes{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r TransactionRoutes) RegisterRouter() {
	handler := handlers.NewTransactionHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	transactionRouters := r.RouteGroup.Group("/transaction")
	transactionRouters.Use(jwt.Use)

	listAdminRouters := transactionRouters.Group("/admin")
	listAdminRouters.Get("", handler.GetListForAdminWithPagination)

	listNormalUserRouters := transactionRouters.Group("/buyer")
	listNormalUserRouters.Get("", handler.GetListForNormalUserWithPagination)

	transactionRouters.Put("/confirm/:id", handler.ConfirmPayment)
	transactionRouters.Put("/cancel/:id", handler.CancelPayment)
	transactionRouters.Get("/:id", handler.GetByID)
	transactionRouters.Post("", handler.Add)
}
