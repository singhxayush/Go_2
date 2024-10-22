package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// app := fiber.New()

	// Custom config
	appConfig := fiber.Config{
		AppName:           "Test App v1.0.1",
		EnablePrintRoutes: true,
		ServerHeader:      "localhost",
		Prefork:       true,

		// CaseSensitive: false,
		// StrictRouting: true,
		// ColorScheme:   fiber.DefaultColors,
	}
	app := fiber.New(appConfig)

	// app.Get("/:name?", func(c *fiber.Ctx) error {
	// 	fmt.Println(c.Params("name"))
	// 	return c.SendString("Hi " + c.Params("name", "id"))
	// })

	// To serve static files such as images, CSS, and JavaScript files, replace your function handler with a file or directory string.
	// app.Static("/", "./public/img.jpg")
	// app.Static("/pirates", "./public/dist")

	app.Get("/", func(c *fiber.Ctx) error {
		log.Info("Received", "req", c.Method(), "path", c.Path())
		return c.Status(http.StatusOK).SendString("Backend is live.")
	})

	app.Get("/items", getAllItems).Name("All Items")
	app.Get("/item/:id", getItem)
	app.Post("/item", createItem)

	log.Info("Starting Server on PORT 8080")

	app.Listen(":8080")
}

// Item represents a product with an ID, title, and price.
// It is used for inventory or item listings.
type Item struct {
	ID    int     `json:"id"`    // Unique identifier for the item
	Title string  `json:"title"` // Title of the item
	Price float32 `json:"price"` // Price of the item
}

var items []Item = []Item{
	{1, "Item 1", 1.5},
	{2, "Item 2", 2.5},
	{3, "Item 3", 3.5},
}

func getAllItems(c *fiber.Ctx) error {
	log.Info("Received", "req", c.Method(), "path", c.Path())

	return c.JSON(items)
}

func getItem(c *fiber.Ctx) error {
	log.Info("Received", "req", c.Method(), "path", c.Path())

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("wrong ID")
	}

	for _, item := range items {
		if item.ID == id {
			return c.JSON(item)
		}
	}

	return c.Status(http.StatusNotFound).SendString("Item not found")

}

func createItem(c *fiber.Ctx) error {
	log.Info("Received", "req", c.Method(), "path", c.Path())

	var item Item

	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
	}

	item.ID = len(items) + 1
	items = append(items, item)

	return c.JSON(item)
}
