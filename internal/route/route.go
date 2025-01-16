package route

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/run-bigpig/conduit/internal/config"
	"github.com/run-bigpig/conduit/internal/handler/auth"
	"github.com/run-bigpig/conduit/internal/handler/copilot"
	"github.com/run-bigpig/conduit/internal/logs"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
)

// Init 加载路由
func Init(ctx context.Context) {
	app := fiber.New(fiber.Config{})
	app.Use(requestid.New())
	app.Use(logs.Logger())
	authRouter(app)
	copilotRouter(app)
	go listen(app)
	log.Info("代理已启动")
	<-ctx.Done()
	if err := app.Shutdown(); err != nil {
		log.Errorf("关闭服务器失败%v", err)
	}
	log.Info("代理已关闭")
}

func listen(app *fiber.App) {
	// 解析证书和私钥
	cert, err := tls.X509KeyPair([]byte(config.ReadString("sslcert")), []byte(config.ReadString("sslkey")))
	if err != nil {
		log.Fatalf("无法解析证书和私钥: %v", err)
	}
	err = app.ListenTLSWithCertificate(fmt.Sprintf(":%d", utils.GetPort()), cert)
	if err != nil {
		log.Errorf("无法启动服务器: %v", err)
	}
}

// 授权路由
func authRouter(app *fiber.App) {
	app.Post("/login/device/code", auth.LoginDeviceCodeHandler)
	app.Get("/login/device", auth.LoginDeviceHandler)
	app.Post("/login/oauth/access_token", auth.LoginOauthAccessTokenHandler)
	app.Get("/copilot_internal/v2/token", auth.InternalTokenHandler)
	app.All("/telemetry", auth.TelemetryHandler)
	app.Get("/user", auth.UserHandler)
	app.Get("/_ping", auth.PingHandler)
	app.Get("/agents", func(ctx *fiber.Ctx) error {
		return utils.Success(ctx, types.M{"agents": []any{}})
	})
	app.Get("/api/v3/meta", auth.ApiV3Handler)
	app.Get("/api/v3/user", auth.UserHandler)
}

// copilot路由
func copilotRouter(app *fiber.App) {
	app.Get("/models", copilot.ModelsHandler)
	app.Post("/v1/engines/copilot-codex/completions", copilot.CodeCompletionHandler)
	app.Post("/v1/engines/copilot-centralus-h100/speculation", copilot.CodeCompletionHandler)
	app.Post("/v1/chat/completions", copilot.ChatCompletionHandler)
	app.Post("/chat/completions", copilot.ChatCompletionHandler)
	app.Post("/embeddings", copilot.EmbeddingsHandler)
}
