package boot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/ecommerce/product-service/domain/configs"
)

type Boot struct {
	App    *fiber.App
	Config *configs.Config
}

func NewBoot(app *fiber.App, config *configs.Config) *Boot {
	return &Boot{
		App:    app,
		Config: config,
	}
}

