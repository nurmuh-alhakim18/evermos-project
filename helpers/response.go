package helpers

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"messsage"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *fiber.Ctx, statusCode int, status bool, message string, errors interface{}, data interface{}) error {
	resp := Response{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	}

	return ctx.Status(statusCode).JSON(resp)
}
