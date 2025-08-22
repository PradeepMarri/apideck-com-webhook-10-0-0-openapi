package main

import (
	"github.com/webhook-api/mcp-server/config"
	"github.com/webhook-api/mcp-server/models"
	tools_webhooks "github.com/webhook-api/mcp-server/tools/webhooks"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_webhooks.CreateWebhooksshortexecuteTool(cfg),
		tools_webhooks.CreateEventlogsallTool(cfg),
		tools_webhooks.CreateWebhookverifyTool(cfg),
		tools_webhooks.CreateWebhooksresolveTool(cfg),
		tools_webhooks.CreateWebhooksallTool(cfg),
		tools_webhooks.CreateWebhooksaddTool(cfg),
		tools_webhooks.CreateWebhooksdeleteTool(cfg),
		tools_webhooks.CreateWebhooksoneTool(cfg),
		tools_webhooks.CreateWebhooksupdateTool(cfg),
		tools_webhooks.CreateWebhooksexecuteTool(cfg),
	}
}
