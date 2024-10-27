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

	// app.Use(addReqID)
	// app.Get("/", reqLoggerMiddleware, getHandler).Name("GET default")
	// app.Get("/item/:id", reqLoggerMiddleware, getItemByID).Name("GET Item")

	app.Get("/", addReqID, reqLoggerMiddleware, getHandler).Name("GET default")
	app.Get("/item/:id", addReqID, reqLoggerMiddleware, getItemByID).Name("GET Item")

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
