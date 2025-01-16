package copilot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/adapter"
	"github.com/run-bigpig/conduit/internal/config"
	"github.com/run-bigpig/conduit/internal/constants"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
)

func CodeCompletionHandler(ctx *fiber.Ctx) error {
	var req types.CopilotCompletionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Fail(ctx, err)
	}
	at := adapter.AdapterFactory(ctx, &types.ApiParams{
		CompletionType: constants.CompletionTypeCode,
		Api:            config.ReadString("code.api"),
		Sk:             config.ReadString("code.sk"),
		Model:          config.ReadString("code.model"),
		Adapter:        config.ReadString("code.adapter"),
	})
	if err := at.Convert(&req); err != nil {
		return utils.Fail(ctx, err)
	}
	return at.Replay()
}
