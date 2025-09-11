package http

import "github.com/gofiber/fiber/v2"

type (
	Handler interface {
		FindAll(c *fiber.Ctx) error
		FindByID(c *fiber.Ctx) error
		Create(c *fiber.Ctx) error
		Update(c *fiber.Ctx) error
		Delete(c *fiber.Ctx) error
	}
)
