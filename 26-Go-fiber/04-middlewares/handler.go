package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func addReqID(c *fiber.Ctx) error {
	uid := uuid.New()
	c.Request().Header.Add("REQUEST-ID", uid.String())
	return c.Next()
}

func reqLoggerMiddleware(c *fiber.Ctx) error {
	reqID := string(c.Request().Header.Peek("REQUEST-ID"))
	log.Info("REQUEST", "method", c.Method(), "path", c.Path(), "reqID", reqID)
	return c.Next()
}

func getHandler(c *fiber.Ctx) error {
	return c.SendString("Backend is live !!!")
}

func getItemByID(c *fiber.Ctx) error {

	item := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "Body lotion for winters",
	}

	return c.JSON(item)

	// res := []byte("Book X")
	// return c.Send(res)
}
