package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetHandler handles requests to the root path ("/") and returns a string
// confirming the backend is live.
//
// Params:
//
//	c - fiber.Ctx (the context for the request)
//
// Returns:
//
//	error - The error if something goes wrong, or nil on success.
func GetHandler(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Backend is live !!!")
}

// Login handles login requests by validating the provided username and password
// against a predefined list of users. If the credentials match, it returns a
// success message; otherwise, it returns a failure message.
//
// Params:
//
//	c - fiber.Ctx (the context for the request, containing request and response objects)
//
// Request Body:
//
//	The request must contain a JSON object with the following fields:
//	  - username (string): The username to authenticate.
//	  - password (string): The password to authenticate.
//
// Returns:
//
//	error - An error if something goes wrong, or nil on success. On success, it sends a JSON response
//	with either a "Login Success" message and a status code 200, or a "Login Failed" message and
//	a status code 401.
func Login(c *fiber.Ctx) error {
	type user struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	// Predefined users for validation
	users := []user{
		{"john_doe", "password123"},
		{"jane_smith", "securepass"},
		{"ayush", "ayush"},
	}

	// Parse request body into a user struct
	var loginDetails user
	if err := c.BodyParser(&loginDetails); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Check if the provided credentials match any user
	for _, u := range users {
		if u.UserName == loginDetails.UserName && u.Password == loginDetails.Password {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login Success"})
		}
	}

	// If credentials don't match
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Login Failed"})
}
