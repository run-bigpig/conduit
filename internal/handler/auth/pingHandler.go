package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
	"time"
)

func PingHandler(ctx *fiber.Ctx) error {
	pong := types.M{
		"now":    time.Now().Second(),
		"status": "ok",
		"ns1":    "200 OK",
	}
	return utils.Success(ctx, pong)
}
