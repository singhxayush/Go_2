package middleware

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddReqID is a middleware function that generates a unique request ID
// for each incoming request and adds it to the request headers. This ID
// can be used for tracking and logging purposes.
//
// Params:
//
//	c - fiber.Ctx (the context for the request, containing request and response objects)
//
// Returns:
//
//	error - An error if something goes wrong, or nil on success. It calls
//	c.Next() to pass control to the next middleware in the chain.
func AddReqID(c *fiber.Ctx) error {
	uid := uuid.New()                                    // Generate a new unique request ID
	c.Request().Header.Add("X-Request-Id", uid.String()) // Add the request ID to the headers
	return c.Next()                                      // Proceed to the next middleware or handler
}

// ReqLoggerMiddleware is a middleware function that logs details of each incoming request,
// including the HTTP method, path, and the associated request ID. This can be useful for
// debugging and monitoring purposes.
//
// Params:
//
//	c - fiber.Ctx (the context for the request, containing request and response objects)
//
// Returns:
//
//	error - An error if something goes wrong, or nil on success. It calls
//	c.Next() to pass control to the next middleware in the chain.
func ReqLoggerMiddleware(c *fiber.Ctx) error {
	reqID := string(c.Request().Header.Peek("X-Request-Id"))                    // Retrieve the request ID from headers
	log.Info("REQUEST", "method", c.Method(), "path", c.Path(), "reqID", reqID) // Log the request details
	return c.Next()                                                             // Proceed to the next middleware or handler
}
