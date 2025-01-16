package logs

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"os"
)

var customLogWriter *LogWriter

func Init(ctx context.Context) {
	customLogWriter = &LogWriter{ctx: ctx}
	iw := io.MultiWriter(customLogWriter, os.Stdout)
	log.SetOutput(iw)
}

func Logger() fiber.Handler {
	return logger.New(logger.Config{Output: customLogWriter})
}

// LogWriter 是一个自定义的 io.Writer，用于捕获日志并发送到前端
type LogWriter struct {
	ctx context.Context
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	msg := string(p)
	runtime.EventsEmit(w.ctx, "log", msg) // 发送日志到前端
	return len(p), nil
}
