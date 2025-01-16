package auth

import (
	"github.com/gofiber/fiber/v2"
)

func LoginDeviceHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("授权成功")
}
