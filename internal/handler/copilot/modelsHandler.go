package copilot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
)

func ModelsHandler(ctx *fiber.Ctx) error {
	models := types.M{
		"data": []types.M{
			{
				"capabilities": types.M{
					"family":    "gpt-3.5-turbo",
					"limits":    types.M{"max_prompt_tokens": 12288, "max_context_window_tokens": 16384, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":                   "gpt-3.5-turbo",
				"name":                 "GPT 3.5 Turbo",
				"object":               "model",
				"version":              "gpt-3.5-turbo-0613",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-3.5-turbo",
					"limits":    types.M{"max_prompt_tokens": 12288, "max_context_window_tokens": 16384, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":                   "gpt-3.5-turbo-0613",
				"name":                 "GPT 3.5 Turbo",
				"object":               "model",
				"version":              "gpt-3.5-turbo-0613",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4",
					"limits":    types.M{"max_prompt_tokens": 32768, "max_context_window_tokens": 32768, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4",
				"name":                 "GPT 4",
				"object":               "model",
				"version":              "gpt-4-0613",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4",
					"limits":    types.M{"max_prompt_tokens": 32768, "max_context_window_tokens": 32768, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4-0613",
				"name":                 "GPT 4",
				"object":               "model",
				"version":              "gpt-4-0613",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4o-mini",
					"limits":    types.M{"max_prompt_tokens": 12288, "max_context_window_tokens": 128000, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"tool_calls": true, "parallel_tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4o-mini",
				"name":                 "GPT 4o Mini",
				"object":               "model",
				"version":              "gpt-4o-mini-2024-07-18",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4o-mini",
					"limits":    types.M{"max_prompt_tokens": 12288, "max_context_window_tokens": 128000, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"tool_calls": true, "parallel_tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4o-mini-2024-07-18",
				"name":                 "GPT 4o Mini",
				"object":               "model",
				"version":              "gpt-4o-mini-2024-07-18",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4-turbo",
					"limits":    types.M{"max_prompt_tokens": 64000, "max_context_window_tokens": 128000, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "cl100k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4-0125-preview",
				"name":                 "GPT 4 Turbo",
				"object":               "model",
				"version":              "gpt-4-0125-preview",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4o",
					"limits":    types.M{"max_prompt_tokens": 64000, "max_context_window_tokens": 128000, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4o",
				"name":                 "GPT 4o",
				"object":               "model",
				"version":              "gpt-4o-2024-05-13",
				"model_picker_enabled": true,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4o",
					"limits":    types.M{"max_prompt_tokens": 64000, "max_context_window_tokens": 128000, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4o-2024-05-13",
				"name":                 "GPT 4o",
				"object":               "model",
				"version":              "gpt-4o-2024-05-13",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4o",
					"limits":    types.M{"max_prompt_tokens": 64000, "max_context_window_tokens": 128000, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4-o-preview",
				"name":                 "GPT 4o",
				"object":               "model",
				"version":              "gpt-4o-2024-05-13",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "gpt-4o",
					"limits":    types.M{"max_prompt_tokens": 64000, "max_context_window_tokens": 128000, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{"parallel_tool_calls": true, "tool_calls": true},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "gpt-4o-2024-08-06",
				"name":                 "GPT 4o",
				"object":               "model",
				"version":              "gpt-4o-2024-08-06",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "o1",
					"limits":    types.M{"max_prompt_tokens": 64000, "max_context_window_tokens": 128000},
					"object":    "model_capabilities",
					"supports":  types.M{},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "o1-preview",
				"name":                 "o1-preview (Preview)",
				"object":               "model",
				"version":              "gpt-4o-2024-08-06",
				"model_picker_enabled": true,
			},
			{
				"capabilities": types.M{
					"family":    "o1",
					"limits":    types.M{"max_prompt_tokens": 64000, "max_context_window_tokens": 128000},
					"object":    "model_capabilities",
					"supports":  types.M{},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "o1-preview-2024-09-12",
				"name":                 "o1-preview (Preview)",
				"object":               "model",
				"version":              "o1-preview-2024-09-12",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "o1-mini",
					"limits":    types.M{"max_prompt_tokens": 20000, "max_context_window_tokens": 128000},
					"object":    "model_capabilities",
					"supports":  types.M{},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "o1-mini",
				"name":                 "o1-mini (Preview)",
				"object":               "model",
				"version":              "o1-preview-2024-09-12",
				"model_picker_enabled": true,
			},
			{
				"capabilities": types.M{
					"family":    "o1-mini",
					"limits":    types.M{"max_prompt_tokens": 20000, "max_context_window_tokens": 128000},
					"object":    "model_capabilities",
					"supports":  types.M{},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "o1-mini-2024-09-12",
				"name":                 "o1-mini (Preview)",
				"object":               "model",
				"version":              "o1-mini-2024-09-12",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "claude-3.5-sonnet",
					"limits":    types.M{"max_prompt_tokens": 195000, "max_context_window_tokens": 200000, "max_output_tokens": 4096},
					"object":    "model_capabilities",
					"supports":  types.M{},
					"tokenizer": "o200k_base",
					"type":      "chat",
				},
				"id":                   "claude-3.5-sonnet",
				"name":                 "Claude 3.5 Sonnet (Preview)",
				"object":               "model",
				"version":              "claude-3.5-sonnet",
				"model_picker_enabled": true,
				"policy": types.M{
					"state": "enabled",
					"terms": "Enable access to the latest Claude 3.5 Sonnet model from Anthropic. [Learn more about how GitHub Copilot serves Claude 3.5 Sonnet](https://docs.github.com/copilot/using-github-copilot/using-claude-sonnet-in-github-copilot).",
				},
			},
			{
				"capabilities": types.M{
					"family":    "text-embedding-ada-002",
					"limits":    types.M{"max_inputs": 256},
					"object":    "model_capabilities",
					"supports":  types.M{},
					"tokenizer": "cl100k_base",
					"type":      "embeddings",
				},
				"id":                   "text-embedding-ada-002",
				"name":                 "Embedding V2 Ada",
				"object":               "model",
				"version":              "text-embedding-ada-002",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "text-embedding-3-small",
					"limits":    types.M{"max_inputs": 256},
					"object":    "model_capabilities",
					"supports":  types.M{"dimensions": true},
					"tokenizer": "cl100k_base",
					"type":      "embeddings",
				},
				"id":                   "text-embedding-3-small",
				"name":                 "Embedding V3 small",
				"object":               "model",
				"version":              "text-embedding-3-small",
				"model_picker_enabled": false,
			},
			{
				"capabilities": types.M{
					"family":    "text-embedding-3-small",
					"object":    "model_capabilities",
					"supports":  types.M{"dimensions": true},
					"tokenizer": "cl100k_base",
					"type":      "embeddings",
				},
				"id":                   "text-embedding-3-small-inference",
				"name":                 "Embedding V3 small (Inference)",
				"object":               "model",
				"version":              "text-embedding-3-small",
				"model_picker_enabled": false,
			},
		},
		"object": "list",
	}
	return utils.Success(ctx, models)
}
