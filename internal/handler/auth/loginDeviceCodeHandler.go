package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/constants"
	"github.com/run-bigpig/conduit/internal/errcode"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
)

func LoginDeviceCodeHandler(ctx *fiber.Ctx) error {
	var req types.LoginDeviceCodeRequest
	if err := ctx.BodyParser(&req); err != nil {
		if err = ctx.QueryParser(&req); err != nil {
			return utils.Fail(ctx, err)
		}
	}
	if req.ClientId == "" {
		return utils.Fail(ctx, errcode.ErrClientIdEmpty)
	}
	code := utils.RandomString(6)
	var authData = types.AuthorizationData{
		ClientId:    req.ClientId,
		UserCode:    code,
		DeviceCode:  utils.RandomString(40),
		DisplayName: "conduit",
	}
	//保存userCode 对应的信息
	authDataStr, _ := json.Marshal(authData)
	if err := utils.SetData(fmt.Sprintf("%s:%s", constants.UserDataHeader, authData.UserCode), authDataStr, 2000); err != nil {
		return utils.Fail(ctx, err)
	}
	//保存deviceCode 对应的userCode
	if err := utils.SetData(fmt.Sprintf("%s:%s", constants.ClientMapUserHeader, authData.DeviceCode), []byte(authData.UserCode), 2000); err != nil {
		return utils.Fail(ctx, err)
	}
	return utils.Success(ctx, types.LoginDeviceCodeResponse{
		DeviceCode:      authData.DeviceCode,
		UserCode:        authData.UserCode,
		VerificationUrl: fmt.Sprintf("https://%s:%d/login/device?user_code=%s", utils.GetDomain(), utils.GetPort(), authData.UserCode),
		ExpiresIn:       2000,
		Interval:        5,
	})
}
