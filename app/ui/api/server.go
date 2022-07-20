package api

import (
	"github.com/davidenq/tweets-timeline-challenge/app/ui/api/handlers"
	_ "github.com/davidenq/tweets-timeline-challenge/docs/api" // docs is generated by Swag CLI, you have to import it.
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func NewServer(port string, handlers handlers.HandlerDefinition) {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/health", handlers.Health)

	apiV1 := app.Group("api/v1")
	apiV1.Get("/users/:username/timeline", handlers.GetTimelineByUser)
	apiV1.Get("/users/:username", handlers.GetUserByUsername)
	apiV1.Put("/users/:username", handlers.UpdateUserByUsername)

	app.Listen(":" + port)
}
