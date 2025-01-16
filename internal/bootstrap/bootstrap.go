package bootstrap

import (
	"context"
	"github.com/run-bigpig/conduit/internal/config"
	"github.com/run-bigpig/conduit/internal/handler/gui"
	"github.com/run-bigpig/conduit/internal/logs"
	"github.com/run-bigpig/conduit/internal/storage"
)

type Boot struct {
	ctx context.Context
}

func NewBoot() *Boot {
	return &Boot{}
}

func (b *Boot) Startup(ctx context.Context) {
	// 初始化配置
	config.Init()
	//初始化日志输出
	logs.Init(ctx)
	//初始化存储
	storage.Init()
	b.ctx = ctx
}

func (b *Boot) NewGuiHandler() *gui.GuiHandler {
	return gui.NewGuiHandler(b.ctx)
}
