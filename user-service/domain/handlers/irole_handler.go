package handlers

import "github.com/gofiber/fiber/v2"

type IRoleHandler interface {
	BrowseAll(ctx *fiber.Ctx) (err error)
}
