package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/ecommerce/product-service/usecases"
)

type HandlerContract struct {
	UseCaseContract *usecases.UseCaseContract
	App             *fiber.App
}
