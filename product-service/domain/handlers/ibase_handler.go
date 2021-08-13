package handlers

import "github.com/gofiber/fiber/v2"

type IBaseHandler interface {
	GetListWithPagination(ctx *fiber.Ctx) (err error)

	GetByID(ctx *fiber.Ctx) (err error)

	Edit(ctx *fiber.Ctx) (err error)

	Add(ctx *fiber.Ctx) (err error)

	Delete(ctx *fiber.Ctx) (err error)
}
