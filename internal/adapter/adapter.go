package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/adapter/deepseek"
	"github.com/run-bigpig/conduit/internal/adapter/openai"
	"github.com/run-bigpig/conduit/internal/types"
)

type Adapter interface {
	Convert(copilotReq *types.CopilotCompletionRequest) error
	Replay() error
}

func AdapterFactory(ctx *fiber.Ctx, params *types.ApiParams) Adapter {
	switch params.Adapter {
	case "openai":
		return openai.New(ctx, params)
	case "deepseek":
		return deepseek.New(ctx, params)
	default:
		return nil
	}
}
