package handlers

import (
	"github.com/davidenq/tweets-timeline-challenge/app/ui/api/httputils/reply"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// ShowAccount   godoc
// @Summary      Get timelines
// @Description  get a list of timelines elements
// @Tags         entities
// @Produce      json
// @Param        username     path      string             true   "twitter's user by username"
// @Param        count        query     string             false  "number of elements"
// @Success      200          {object}  api.Response
// @Failure      400,404,500  {object}  api.ErrorResponse
// @Router       /users/:username/timeline [get]
func (h *Base) GetTimelineByUser(c *fiber.Ctx) error {
	username := c.Params("username")
	count := c.Query("count")

	if count == "" {
		count = "5"
	}
	timelines, err := h.timelineUsecase.GetTimeline(c.Context(), username, count)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return reply.Send(c, err)
	}
	return reply.Send(c, timelines)
}
