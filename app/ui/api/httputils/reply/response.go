package reply

import (
	"encoding/json"
	"net/http"

	"github.com/davidenq/tweets-timeline-challenge/app/utils/errorkit"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func Send(ctx *fiber.Ctx, data interface{}) error {
	var res Response
	var errRes ErrorResponse
	switch val := data.(type) {
	case *errorkit.AppError:
		errRes.Message = val.Error()
		errRes.Error = http.StatusText(val.Code)
		if val.Code == 0 {
			res.Code = http.StatusInternalServerError
		}
		errRes.Code = val.Code
		return ctx.Status(errRes.Code).JSON(errRes)
	default:
		res.Data = data
		message, _ := json.Marshal(res)
		log.Info().Msg(string(message))
		return ctx.Status(200).JSON(res)
	}

}
