package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Empty(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Server started,request to api/{controller}/{action}",
	})
}
