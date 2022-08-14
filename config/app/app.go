package app

import (
	"fmt"
	"time"
	"web-api/config/env"
	"web-api/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Application() (app *fiber.App, server string) {
	initTimeZone()
	server = initServer()
	app = initFiber()
	route.SetupRoutes(app)
	return app, server
}

func initFiber() (app *fiber.App) {
	app = fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))
	return app
}

func initServer() (server string) {
	e := env.Network()
	server = fmt.Sprint(e.HOST + ":" + e.PORT)
	return server
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

