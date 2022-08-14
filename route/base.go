package route

import (
	"github.com/gofiber/fiber/v2"
	"web-api/controller"
)

func SetupRoutes(router *fiber.App) {
	router.Get("/", controller.Empty)

	api := router.Group("/api")
	{
		SuRoutes(api)
	}
}
