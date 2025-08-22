package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/webhook-api/mcp-server/config"
	"github.com/webhook-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func WebhooksresolveHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		serviceIdVal, ok := args["serviceId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: serviceId"), nil
		}
		serviceId, ok := serviceIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: serviceId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["e"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("e=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/webhook/w/%s/%s%s", cfg.BaseURL, id, serviceId, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Authorization", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.ResolveWebhookResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateWebhooksresolveTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_webhook_w_id_serviceId",
		mcp.WithDescription("Resolve and Execute a connection webhook"),
		mcp.WithString("id", mcp.Required(), mcp.Description("JWT Webhook token that represents the connection lookupId. Signed so we know source came from us")),
		mcp.WithString("serviceId", mcp.Required(), mcp.Description("Service provider ID.")),
		mcp.WithString("e", mcp.Description("The name of downstream event when connector does not supply in body or header")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    WebhooksresolveHandler(cfg),
	}
}
