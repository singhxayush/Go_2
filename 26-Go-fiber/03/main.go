package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig := fiber.Config{
		EnablePrintRoutes: true,
		ServerHeader:      "Awesome App 1",
		Immutable:         true,
	}
	app := fiber.New(appConfig)

	app.Get("/", getHandler).Name("GET default")
	app.Get("/books", getAllBooks).Name("GET Books")

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getHandler(c *fiber.Ctx) error {
	log.Info("request", "method", c.Method(), "path", c.Path())
	return c.SendString("Backend is live !!!")
}
