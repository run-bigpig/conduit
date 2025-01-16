package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
	"time"
)

func InternalTokenHandler(ctx *fiber.Ctx) error {
	fakeData, err := copilotFakeData()
	if err != nil {
		return utils.Fail(ctx, err)
	}
	return utils.Success(ctx, fakeData)
}

func copilotFakeData() (map[string]interface{}, error) {
	trackingId, _ := uuid.NewUUID()
	sku := "copilot_for_business_seat"
	copilotToken, err := utils.CreateJwt(map[string]interface{}{
		"tid":  trackingId,
		"sku":  sku,
		"st":   "dotcom",
		"chat": 1,
		"u":    "github",
	}, 1800)
	if err != nil {
		return nil, err
	}
	endpoints := make(map[string]interface{})
	endpoints["api"] = fmt.Sprintf("https://%s.%s:%d", "api", utils.GetDomain(), utils.GetPort())
	endpoints["origin-tracker"] = "https://origin-tracker.individual.githubcopilot.com"
	endpoints["proxy"] = fmt.Sprintf("https://%s.%s:%d", "copilot-proxy", utils.GetDomain(), utils.GetPort())
	endpoints["telemetry"] = fmt.Sprintf("https://%s.%s:%d", "copilot-telemetry-service", utils.GetDomain(), utils.GetPort())

	return types.M{
		"annotations_enabled":                      true,
		"chat_enabled":                             true,
		"chat_jetbrains_enabled":                   true,
		"code_quote_enabled":                       true,
		"code_review_enabled":                      false,
		"codesearch":                               true,
		"copilot_ide_agent_chat_gpt4_small_prompt": false,
		"copilotignore_enabled":                    false,
		"endpoints":                                endpoints,
		"expires_at":                               time.Now().Unix() + 1800,
		"individual":                               true,
		"nes_enabled":                              false,
		"prompt_8k":                                true,
		"public_suggestions":                       "disabled",
		"refresh_in":                               1500,
		"sku":                                      sku,
		"snippy_load_test_enabled":                 false,
		"telemetry":                                "disabled",
		"token":                                    copilotToken,
		"tracking_id":                              trackingId,
		"intellij_editor_fetcher":                  false,
		"vsc_electron_fetcher":                     false,
		"vs_editor_fetcher":                        false,
		"vsc_panel_v2":                             false,
		"xcode":                                    true,
		"xcode_chat":                               true,
		"limited_user_quotas":                      nil,
		"limited_user_reset_date":                  nil,
		"vsc_electron_fetcher_v2":                  false,
	}, nil
}
