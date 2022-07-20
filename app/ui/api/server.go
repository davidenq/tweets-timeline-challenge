package api

import (
	"github.com/davidenq/tweets-timeline-challenge/app/ui/api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	//swagger_doc_import //this line must not removed
)

func NewServer(port string, handlers handlers.HandlerDefinition) {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/health", handlers.Health)

	apiV1 := app.Group("api/v1")
	apiV1.Get("/users/:username/timeline", handlers.GetTimelineByUser)
	apiV1.Get("/users/:username", handlers.GetUserByUsername)
	apiV1.Put("/users/:username", handlers.UpdateUserByUsername)

	app.Listen(":" + port)
}
