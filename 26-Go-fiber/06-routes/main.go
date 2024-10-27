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

	app.Use(requestid.New(), logger.New()) // Go fiber internal middlewares

	api := app.Group("/api")

	// app.Use(middleware.AddReqID, middleware.ReqLoggerMiddleware) // User made middlewares

	v1 := api.Group("/v1")

	v1.Get("/", handler.GetHandler).Name("Root")
	v1.Post("/login", handler.Login).Name("User Login")
	v1.Get("/metrics", monitor.New()).Name("Metrics")

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
