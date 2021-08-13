package handlers

import "github.com/gofiber/fiber/v2"

type ICategoryHandler interface {
	IBaseHandler
	GetAll(ctx *fiber.Ctx) (err error)
}
