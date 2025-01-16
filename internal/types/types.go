package types

import "github.com/run-bigpig/conduit/internal/constants"

type M map[string]interface{}

type LoginDeviceCodeRequest struct {
	ClientId string `json:"client_id" query:"client_id"` // 客户端id
	Scope    string `json:"scope" query:"scope"`
}

type LoginDeviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`      // 设备代码
	UserCode        string `json:"user_code"`        // 用户代码
	VerificationUrl string `json:"verification_uri"` // 验证地址
	ExpiresIn       int    `json:"expires_in"`       // 过期时间
	Interval        int    `json:"interval"`         // 间隔时间
}

type LoginOauthAccessTokenRequest struct {
	DeviceCode string `json:"device_code" query:"device_code"`
	ClientId   string `json:"client_id" query:"client_id"`
	GrantType  string `json:"grant_type" query:"grant_type"`
}

type AuthorizationData struct {
	ClientId    string `json:"client_id"`
	UserCode    string `json:"user_code"`
	DeviceCode  string `json:"device_code"`
	CardCode    string `json:"card_code"`
	DisplayName string `json:"display_name"`
}

type ApiParams struct {
	CompletionType constants.CompletionType
	Api            string
	Sk             string
	Model          string
	Adapter        string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CopilotCompletionRequest struct {
	Model        string     `json:"model"`
	Prompt       string     `json:"prompt,omitempty"`
	Suffix       string     `json:"suffix,omitempty"`
	Temperature  float64    `json:"temperature"`
	TopP         float64    `json:"top_p"`
	N            int        `json:"n,omitempty"`
	Stream       bool       `json:"stream"`
	Stop         []string   `json:"stop"`
	MaxTokens    int        `json:"max_tokens"`
	Messages     []*Message `json:"messages,omitempty"`
	Tools        []*Tool    `json:"tools,omitempty"`
	ToolChoice   any        `json:"tool_choice,omitempty"`
	FunctionCall any        `json:"function_call,omitempty"`
	Functions    any        `json:"functions,omitempty"`
}

type Tool struct {
	Id       string    `json:"id,omitempty"`
	Type     string    `json:"type,omitempty"`
	Function *Function `json:"function"`
}

type Function struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Parameters  any    `json:"parameters,omitempty"`
	Arguments   any    `json:"arguments,omitempty"`
}

type Hosts struct {
	Ip     string
	Domain string
}

type SystemConfigRequest struct {
	Domain  string `json:"domain"`
	Port    int    `json:"port"`
	SslCert string `json:"sslCert"`
	SslKey  string `json:"sslKey"`
}

type CompletionConfigRequest struct {
	Api       string `json:"api"`
	Sk        string `json:"sk"`
	Model     string `json:"model"`
	Adapter   string `json:"adapter"`
	MaxTokens int    `json:"maxTokens"`
}

type GuiResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type GuiAllConfig struct {
	Domain  string                   `json:"domain"`
	Port    int                      `json:"port"`
	SslCert string                   `json:"sslCert"`
	SslKey  string                   `json:"sslKey"`
	Chat    *CompletionConfigRequest `json:"chat"`
	Code    *CompletionConfigRequest `json:"code"`
}
