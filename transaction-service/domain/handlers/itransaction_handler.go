package handlers

import "github.com/gofiber/fiber/v2"

type ITransactionHandler interface {
	GetListForAdminWithPagination(ctx *fiber.Ctx) (err error)

	GetListForNormalUserWithPagination(ctx *fiber.Ctx) (err error)

	GetByID(ctx *fiber.Ctx) (err error)

	Add(ctx *fiber.Ctx) (err error)

	CancelPayment(ctx *fiber.Ctx) (err error)

	ConfirmPayment(ctx *fiber.Ctx) (err error)
}
