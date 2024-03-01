package routers

import (
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App
}

func (c *RouteConfig) Setup() {
	c.App.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
}
