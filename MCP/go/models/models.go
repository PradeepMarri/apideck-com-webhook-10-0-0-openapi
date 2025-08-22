package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// CreateWebhookRequest represents the CreateWebhookRequest schema from the OpenAPI specification
type CreateWebhookRequest struct {
	Events []string `json:"events"` // The list of subscribed events for this webhook. [`*`] indicates that all events are enabled.
	Status string `json:"status"` // The status of the webhook.
	Unified_api string `json:"unified_api"` // Name of Apideck Unified API
	Delivery_url string `json:"delivery_url"` // The delivery url of the webhook endpoint.
	Description string `json:"description,omitempty"` // A description of the object.
}

// UpdateWebhookResponse represents the UpdateWebhookResponse schema from the OpenAPI specification
type UpdateWebhookResponse struct {
	Data Webhook `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// WebhookEvent represents the WebhookEvent schema from the OpenAPI specification
type WebhookEvent struct {
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
	Entity_url string `json:"entity_url,omitempty"` // The url to retrieve entity detail.
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Event_type string `json:"event_type,omitempty"`
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
}

// ExecuteWebhookEventRequest represents the ExecuteWebhookEventRequest schema from the OpenAPI specification
type ExecuteWebhookEventRequest struct {
}

// GetWebhooksResponse represents the GetWebhooksResponse schema from the OpenAPI specification
type GetWebhooksResponse struct {
	Data []Webhook `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// GetWebhookResponse represents the GetWebhookResponse schema from the OpenAPI specification
type GetWebhookResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Webhook `json:"data"`
	Status string `json:"status"` // HTTP Response Status
}

// WebhookEventLog represents the WebhookEventLog schema from the OpenAPI specification
type WebhookEventLog struct {
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // Number of attempts webhook endpoint was called before a success was returned or eventually failed
	Status_code int `json:"status_code,omitempty"` // HTTP Status code that was returned.
	Http_method string `json:"http_method,omitempty"` // HTTP Method of request to endpoint.
	Id string `json:"id,omitempty"`
	Endpoint string `json:"endpoint,omitempty"` // The URL of the webhook endpoint.
	Entity_type string `json:"entity_type,omitempty"` // Name of the Entity described by the attributes delivered within payload
	Consumer_id string `json:"consumer_id,omitempty"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Request_body string `json:"request_body,omitempty"` // The JSON stringified payload that was delivered to the webhook endpoint.
	Event_type string `json:"event_type,omitempty"` // Name of source event that webhook is subscribed to.
	Retry_scheduled bool `json:"retry_scheduled,omitempty"` // If the request has not hit the max retry limit and will be retried.
	Response_body string `json:"response_body,omitempty"` // The JSON stringified response that was returned from the webhook endpoint.
	Service map[string]interface{} `json:"service,omitempty"` // Apideck service provider associated with event.
	Application_id string `json:"application_id,omitempty"` // ID of your Apideck Application
	Success bool `json:"success,omitempty"` // Whether or not the request was successful.
	Timestamp string `json:"timestamp,omitempty"` // ISO Date and time when the request was made.
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Attempts []map[string]interface{} `json:"attempts,omitempty"` // record of each attempt to call webhook endpoint
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
}

// DeleteWebhookResponse represents the DeleteWebhookResponse schema from the OpenAPI specification
type DeleteWebhookResponse struct {
	Data Webhook `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// CreateWebhookResponse represents the CreateWebhookResponse schema from the OpenAPI specification
type CreateWebhookResponse struct {
	Data Webhook `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// ExecuteWebhookResponse represents the ExecuteWebhookResponse schema from the OpenAPI specification
type ExecuteWebhookResponse struct {
	Request_id string `json:"request_id,omitempty"` // UUID of the request received
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Timestamp string `json:"timestamp,omitempty"` // ISO Datetime webhook event was received
}

// WebhookEventLogsFilter represents the WebhookEventLogsFilter schema from the OpenAPI specification
type WebhookEventLogsFilter struct {
	Exclude_apis string `json:"exclude_apis,omitempty"`
	Service map[string]interface{} `json:"service,omitempty"`
	Consumer_id string `json:"consumer_id,omitempty"`
	Entity_type string `json:"entity_type,omitempty"`
	Event_type string `json:"event_type,omitempty"`
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// ResolveWebhookEventRequest represents the ResolveWebhookEventRequest schema from the OpenAPI specification
type ResolveWebhookEventRequest struct {
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// UpdateWebhookRequest represents the UpdateWebhookRequest schema from the OpenAPI specification
type UpdateWebhookRequest struct {
	Description string `json:"description,omitempty"` // A description of the object.
	Events []string `json:"events,omitempty"` // The list of subscribed events for this webhook. [`*`] indicates that all events are enabled.
	Status string `json:"status,omitempty"` // The status of the webhook.
	Delivery_url string `json:"delivery_url,omitempty"` // The delivery url of the webhook endpoint.
}

// ResolveWebhookResponse represents the ResolveWebhookResponse schema from the OpenAPI specification
type ResolveWebhookResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Timestamp string `json:"timestamp,omitempty"` // ISO Datetime webhook event was received
	Request_id string `json:"request_id,omitempty"` // UUID of the request received
}

// Webhook represents the Webhook schema from the OpenAPI specification
type Webhook struct {
	Execute_base_url string `json:"execute_base_url"` // The Unify Base URL events from connectors will be sent to after service id is appended.
	Id string `json:"id,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Status string `json:"status"` // The status of the webhook.
	Unified_api string `json:"unified_api"` // Name of Apideck Unified API
	Disabled_reason string `json:"disabled_reason,omitempty"` // Indicates if the webhook has has been disabled as it reached its retry limit or if account is over the usage allocated by it's plan.
	Description string `json:"description,omitempty"` // A description of the object.
	Delivery_url string `json:"delivery_url"` // The delivery url of the webhook endpoint.
	Events []string `json:"events"` // The list of subscribed events for this webhook. [`*`] indicates that all events are enabled.
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// GetWebhookEventLogsResponse represents the GetWebhookEventLogsResponse schema from the OpenAPI specification
type GetWebhookEventLogsResponse struct {
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []WebhookEventLog `json:"data"`
}
