package handlers

import "github.com/gofiber/fiber/v2"

type IFileHandler interface {
	Upload(ctx *fiber.Ctx) (err error)

	GetUrlByKey(ctx *fiber.Ctx) (err error)
}
