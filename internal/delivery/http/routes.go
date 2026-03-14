package fiberhandler

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, ingredientHandler *IngredientHandler) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "service running",
		})
	})

	api := app.Group("/api/v1")

	api.Get("/ingredients", ingredientHandler.List)
	api.Post("/ingredients", ingredientHandler.Create)
	api.Put("/ingredients/:uuid", ingredientHandler.Update)
	api.Delete("/ingredients/:uuid", ingredientHandler.Delete)
}
