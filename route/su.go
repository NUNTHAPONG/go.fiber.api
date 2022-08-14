package route

import (
	"github.com/gofiber/fiber/v2"
	"web-api/controller/su"
)

func SuRoutes(router fiber.Router) {
	api := router.Group("/su")
	{
		api.Get("/param", su.GetParameter)
		api.Get("/page", su.ListPage)
		api.Put("/page", su.UpdatePage)
		api.Post("/page", su.CreatePage)
		api.Delete("/page", su.DeletePage)
	}
}
