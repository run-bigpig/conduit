package openai

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
	"io"
)

type OpenAi struct {
	ctx    *fiber.Ctx
	params *types.ApiParams
	body   []byte
}

func New(ctx *fiber.Ctx, params *types.ApiParams) *OpenAi {
	return &OpenAi{
		ctx:    ctx,
		params: params,
		body:   make([]byte, 0),
	}
}

func (o *OpenAi) Convert(req *types.CopilotCompletionRequest) error {
	req.Model = o.params.Model
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	o.body = data
	return nil
}

func (o *OpenAi) Replay() error {
	header := map[string]any{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.params.Sk,
	}
	resp, err := utils.SendRequest(o.params.Api, fiber.MethodPost, header, o.body)
	if err != nil {
		log.Errorf("Error while sending request: %s. Closing http connection.", err.Error())
		return err
	}
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		log.Errorf("Error while sending request: %s. Closing http connection.", string(body))
		return utils.CompletionAbort(o.ctx, resp.StatusCode)
	}
	return utils.StreamResponse(o.ctx, resp)
}
