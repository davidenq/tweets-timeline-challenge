package handlers

import (
	"github.com/davidenq/tweets-timeline-challenge/app/ui/api/httputils/reply"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// ShowAccount   godoc
// @Summary      Get User
// @Description  It is used to get a list of entities such as user or oauth (this last one it is just for testing purpose)
// @Tags         entities
// @Produce      json
// @Param        entityName  path      string  true   "name entity"              Enums(user, oauth)
// @Param        id          query     string  false  "search entity by id"
// @Param        name        query     string  false  "search entities by name"
// @Success      200  {object}  api.Response
// @Failure 400,404,500 {object} api.ErrorResponse
// @Router       /user/:username [get]
func (h *Base) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	err := h.userUsecase.Update(c.Context(), username)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return reply.Send(c, err)
	}
	return reply.Send(c, "")
}

// ShowAccount   godoc
// @Summary      Update User
// @Description  It is used to get a list of entities such as user or oauth (this last one it is just for testing purpose)
// @Tags         entities
// @Produce      json
// @Param        entityName  path      string  true   "name entity"              Enums(user, oauth)
// @Param        id          query     string  false  "search entity by id"
// @Param        name        query     string  false  "search entities by name"
// @Success      200  {object}  api.Response
// @Failure 400,404,500 {object} api.ErrorResponse
// @Router       /user/:username[put]
func (h *Base) UpdateUserByUsername(c *fiber.Ctx) error {
	return nil
}
