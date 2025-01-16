package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/conduit/internal/constants"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
)

func LoginOauthAccessTokenHandler(ctx *fiber.Ctx) error {
	var req types.LoginOauthAccessTokenRequest
	if err := ctx.BodyParser(&req); err != nil {
		if err = ctx.QueryParser(&req); err != nil {
			return utils.Fail(ctx, err)
		}
	}
	//获取设备码映射用户码
	userCode, err := utils.GetData(fmt.Sprintf("%s:%s", constants.ClientMapUserHeader, req.DeviceCode))
	if err != nil {
		return utils.Fail(ctx, err)
	}
	//获取用户码映射用户信息
	userInfo, err := utils.GetData(fmt.Sprintf("%s:%s", constants.UserDataHeader, userCode))
	if err != nil {
		return utils.Fail(ctx, err)
	}
	var authData types.AuthorizationData
	err = json.Unmarshal(userInfo, &authData)
	if err != nil {
		return utils.Fail(ctx, err)
	}
	claims := map[string]interface{}{
		"userCode":    userCode,
		"deviceCode":  req.DeviceCode,
		"displayName": authData.DisplayName,
		"clientId":    authData.ClientId,
	}
	token, err := utils.CreateJwt(claims, 1800)
	if err != nil {
		return utils.Fail(ctx, err)
	}
	return utils.Success(ctx, map[string]interface{}{
		"access_token": token,
		"scope":        "",
		"token_type":   "bearer",
	})
}
