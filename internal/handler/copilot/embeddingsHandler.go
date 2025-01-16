package copilot

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
)

func EmbeddingsHandler(ctx *fiber.Ctx) error {
	return utils.Json(ctx, 500, types.M{"error": errors.New("Not implemented.")})
}
