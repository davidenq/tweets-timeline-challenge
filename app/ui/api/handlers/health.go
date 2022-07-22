package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ShowAccount   godoc
// @Summary      Check healthy of application
// @Description  It is used to check the status of the current API Rest
// @Tags         health
// @Produce      plain
// @Success      200  {string}  string "ok"
// @Router       /health [get]
func (h *Base) Health(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
