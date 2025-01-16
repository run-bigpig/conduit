package gui

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/run-bigpig/conduit/internal/config"
	"github.com/run-bigpig/conduit/internal/errcode"
	"github.com/run-bigpig/conduit/internal/route"
	"github.com/run-bigpig/conduit/internal/system"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
)

var (
	running bool
)

type GuiHandler struct {
	ctx          context.Context
	appCtx       context.Context
	appCtxCancel context.CancelFunc
}

func NewGuiHandler(ctx context.Context) *GuiHandler {
	return &GuiHandler{ctx: ctx}
}

func (g *GuiHandler) GetStatus() *types.GuiResponse {
	return utils.GuiSuccess(running)
}

func (g *GuiHandler) Start() *types.GuiResponse {
	if !g.checkConfig() {
		return utils.GuiFail(1, errcode.ErrInvalidParams)
	}
	g.appCtx, g.appCtxCancel = context.WithCancel(context.Background())
	go route.Init(g.appCtx)
	running = true
	return utils.GuiSuccess(nil)
}

func (g *GuiHandler) Stop() *types.GuiResponse {
	g.appCtxCancel()
	running = false
	return utils.GuiSuccess(nil)
}

func (g *GuiHandler) checkConfig() bool {
	conf, ok := g.GetConfig().Data.(*types.GuiAllConfig)
	if !ok {
		return false
	}
	return conf.Port > 0 && conf.Domain != "" && conf.SslCert != "" && conf.SslKey != "" && conf.Chat.Api != "" && conf.Code.Api != "" && conf.Chat.Sk != "" && conf.Code.Sk != "" && conf.Code.MaxTokens > 0 && conf.Chat.MaxTokens > 0 && conf.Chat.Adapter != "" && conf.Code.Adapter != "" && conf.Chat.Model != "" && conf.Code.Model != ""
}

func (g *GuiHandler) GetConfig() *types.GuiResponse {
	conf := &types.GuiAllConfig{
		Chat: &types.CompletionConfigRequest{
			Adapter:   config.ReadString("chat.adapter"),
			Api:       config.ReadString("chat.api"),
			Model:     config.ReadString("chat.model"),
			Sk:        config.ReadString("chat.sk"),
			MaxTokens: config.ReadInt("chat.maxtokens"),
		},
		Code: &types.CompletionConfigRequest{
			Adapter:   config.ReadString("code.adapter"),
			Api:       config.ReadString("code.api"),
			Model:     config.ReadString("code.model"),
			Sk:        config.ReadString("code.sk"),
			MaxTokens: config.ReadInt("code.maxtokens"),
		},
		Domain:  config.ReadString("domain"),
		Port:    config.ReadInt("port"),
		SslCert: config.ReadString("sslcert"),
		SslKey:  config.ReadString("sslkey"),
	}
	return utils.GuiSuccess(conf)
}

// SetSystemConfig 设置系统配置
func (g *GuiHandler) SetSystemConfig(req *types.SystemConfigRequest) *types.GuiResponse {
	if req.Domain == "" || req.Port == 0 || req.SslCert == "" || req.SslKey == "" {
		return utils.GuiFail(1, errcode.ErrInvalidParams)
	}
	domainList := []string{req.Domain, fmt.Sprintf("*.%s", req.Domain)}
	if !utils.ParseCertificateDomain(req.SslCert, domainList) {
		return utils.GuiFail(1, errcode.ErrCertInvalid)
	}
	err := system.SyncHost()
	if err != nil {
		return utils.GuiFail(1, err)
	}
	config.Write("domain", req.Domain, true)
	config.Write("port", req.Port, true)
	config.Write("sslcert", req.SslCert, true)
	config.Write("sslkey", req.SslKey, true)
	return utils.GuiSuccess(nil)
}

// SetCodeCompletionConfig 设置代码补全配置
func (g *GuiHandler) SetCodeCompletionConfig(req *types.CompletionConfigRequest) *types.GuiResponse {
	log.Infof("Set code completion config: %+v", req)
	if req.Adapter == "" || req.Api == "" || req.Model == "" || req.Sk == "" || req.MaxTokens == 0 {
		return utils.GuiFail(1, errcode.ErrInvalidParams)
	}
	config.Write("code.adapter", req.Adapter, true)
	config.Write("code.api", req.Api, true)
	config.Write("code.model", req.Model, true)
	config.Write("code.sk", req.Sk, true)
	config.Write("code.maxtokens", req.MaxTokens, true)
	return utils.GuiSuccess(nil)
}

// SetChatCompletionConfig 设置对话补全配置
func (g *GuiHandler) SetChatCompletionConfig(req *types.CompletionConfigRequest) *types.GuiResponse {
	if req.Adapter == "" || req.Api == "" || req.Model == "" || req.Sk == "" || req.MaxTokens == 0 {
		return utils.GuiFail(1, errcode.ErrInvalidParams)
	}
	config.Write("chat.adapter", req.Adapter, true)
	config.Write("chat.api", req.Api, true)
	config.Write("chat.model", req.Model, true)
	config.Write("chat.sk", req.Sk, true)
	config.Write("chat.maxtokens", req.MaxTokens, true)
	return utils.GuiSuccess(nil)

}
