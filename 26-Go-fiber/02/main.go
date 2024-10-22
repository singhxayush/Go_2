package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Custom config
	appConfig := fiber.Config{
		AppName:           "Test App v1.0.1",
		EnablePrintRoutes: true,
		ServerHeader:      "Awesome App 1",
		Immutable: true,
	}
	app := fiber.New(appConfig)

	app.Get("/:name", getHandler).Name("GET default")

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

const gridID = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var idx = 0

func getHandlerID() string {
	c := gridID[idx%26]
	idx++
	return fmt.Sprintf("gridID-%v-%c", idx, c)
}

func getHandler(c *fiber.Ctx) error {
	ccID := getHandlerID()
	n := c.Params("name")

	go func() {
		slog.Info("handler start", "ccID", ccID, "name", n)
		t := time.After(10 * time.Second)
		for {
			select {
			case <-t:
				slog.Info("handler finish", "ccID", ccID, "name", n)
				return
			default:
				slog.Info("handler running", "ccID", ccID, "name", n)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// slog.Info("Request Params", "name", n)
	return nil
}
