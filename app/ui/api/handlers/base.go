package handlers

import (
	"github.com/davidenq/tweets-timeline-challenge/app/domain"
	"github.com/gofiber/fiber/v2"
)

type HandlerDefinition interface {
	Health(c *fiber.Ctx) error
	GetTimelineByUser(c *fiber.Ctx) error
	GetUserByUsername(c *fiber.Ctx) error
	UpdateUserByUsername(c *fiber.Ctx) error
}

type Base struct {
	timelineUsecase domain.TimeLineUsecaseFacadeDrivingPort
	userUsecase     domain.UserUsecaseDrivingPort
}

func NewHandlers(
	timelineUsecase domain.TimeLineUsecaseFacadeDrivingPort,
	userUsecase domain.UserUsecaseDrivingPort,
) HandlerDefinition {
	return &Base{
		timelineUsecase: timelineUsecase,
		userUsecase:     userUsecase,
	}
}
