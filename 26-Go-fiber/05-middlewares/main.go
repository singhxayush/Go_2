package main

import (
	"fiber/handler"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	appConfig := fiber.Config{
		EnablePrintRoutes: true,
		ServerHeader:      "Awesome App 1",
		// Immutable:         true,
	}
	app := fiber.New(appConfig)

	// app.Use(middleware.AddReqID, middleware.ReqLoggerMiddleware) // User made middlewares
	app.Use(requestid.New(), logger.New()) // Go fiber internal middlewares

	app.Get("/", handler.GetHandler).Name("Root")
	app.Post("/login", handler.Login).Name("User Login")
	app.Get("/metrics", monitor.New())

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
