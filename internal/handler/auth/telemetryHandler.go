package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
)

func TelemetryHandler(ctx *fiber.Ctx) error {
	if ctx.Method() == fiber.MethodGet {
		return ctx.Status(fiber.StatusNotFound).SendString("Not Found")
	}
	var jsonData []interface{}
	telemetryData := types.M{
		"itemsReceived": 0,
		"itemsAccepted": 0,
		"appId":         nil,
		"errors":        []string{},
	}
	if err := ctx.BodyParser(&jsonData); err != nil {
		return utils.Success(ctx, telemetryData)
	}
	telemetryData["itemsReceived"] = len(jsonData)
	telemetryData["itemsAccepted"] = len(jsonData)
	return utils.Success(ctx, telemetryData)
}
