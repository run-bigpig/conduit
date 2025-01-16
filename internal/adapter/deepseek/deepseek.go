package deepseek

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/adapter/openai"
	"github.com/run-bigpig/conduit/internal/config"
	"github.com/run-bigpig/conduit/internal/constants"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/spf13/cast"
)

type DeepSeek struct {
	ctx    *fiber.Ctx
	oai    *openai.OpenAi
	params *types.ApiParams
}

func New(ctx *fiber.Ctx, params *types.ApiParams) *DeepSeek {
	return &DeepSeek{
		ctx:    ctx,
		params: params,
		oai:    openai.New(ctx, params),
	}
}

func (o *DeepSeek) Convert(req *types.CopilotCompletionRequest) error {
	switch o.params.CompletionType {
	case constants.CompletionTypeCode:
		req.N = 1
		req.TopP = 1.0
		req.Temperature = 0
		req.MaxTokens = cast.ToInt(config.Read("code.maxtokens"))
	case constants.CompletionTypeChat:
		req.Tools = nil
		req.ToolChoice = nil
		req.Functions = nil
		req.FunctionCall = nil
		req.Temperature = 1.3
		req.MaxTokens = cast.ToInt(config.Read("chat.maxtokens"))
	}
	return o.oai.Convert(req)
}

func (o *DeepSeek) Replay() error {
	return o.oai.Replay()
}
