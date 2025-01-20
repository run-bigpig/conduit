package copilot

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/adapter"
	"github.com/run-bigpig/conduit/internal/config"
	"github.com/run-bigpig/conduit/internal/constants"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
	"strings"
)

func ChatCompletionHandler(ctx *fiber.Ctx) error {
	var req types.CopilotCompletionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Fail(ctx, err)
	}
	at := adapter.AdapterFactory(ctx, &types.ApiParams{
		CompletionType: constants.CompletionTypeChat,
		Api:            config.ReadString("chat.api"),
		Sk:             config.ReadString("chat.sk"),
		Model:          config.ReadString("chat.model"),
		Adapter:        config.ReadString("chat.adapter"),
	})
	for _, msg := range req.Messages {
		if msg.Role == "system" {
			msg.Content = strings.ReplaceAll(msg.Content, "locale: en", fmt.Sprintf("locale: %s", config.ReadString("chat.locale")))
		}
	}
	if err := at.Convert(&req); err != nil {
		return utils.Fail(ctx, err)
	}
	return at.Replay()
}
