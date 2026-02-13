// PicoClaw - Ultra-lightweight personal AI agent
// Inspired by and based on nanobot: https://github.com/HKUDS/nanobot
// License: MIT
//
// Copyright (c) 2026 PicoClaw contributors

package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/weiwei929/mypicoclaw/pkg/config"
	"github.com/weiwei929/mypicoclaw/pkg/logger"
)

type HTTPProvider struct {
	apiKey     string
	apiBase    string
	httpClient *http.Client
	// Fallback provider for disaster recovery
	fallbackKey   string
	fallbackBase  string
	fallbackModel string
}

func NewHTTPProvider(apiKey, apiBase string) *HTTPProvider {
	return &HTTPProvider{
		apiKey:  apiKey,
		apiBase: apiBase,
		httpClient: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

// NewHTTPProviderWithFallback creates a provider with automatic failover to a backup model.
func NewHTTPProviderWithFallback(apiKey, apiBase, fallbackKey, fallbackBase, fallbackModel string) *HTTPProvider {
	p := NewHTTPProvider(apiKey, apiBase)
	p.fallbackKey = fallbackKey
	p.fallbackBase = fallbackBase
	p.fallbackModel = fallbackModel
	return p
}

func (p *HTTPProvider) Chat(ctx context.Context, messages []Message, tools []ToolDefinition, model string, options map[string]interface{}) (*LLMResponse, error) {
	if p.apiBase == "" {
		return nil, fmt.Errorf("API base not configured")
	}

	// Try primary model first
	resp, err := p.chatDirect(ctx, messages, tools, model, options, p.apiKey, p.apiBase)
	if err == nil {
		return resp, nil
	}

	// If fallback is configured and error is not a client-side issue, try fallback
	if p.fallbackBase != "" && p.fallbackModel != "" && p.isFailoverEligible(err) {
		logger.WarnCF("provider", fmt.Sprintf("⚡ Primary model failed, switching to fallback: %s", p.fallbackModel),
			map[string]interface{}{
				"primary_error":  err.Error(),
				"fallback_model": p.fallbackModel,
			})
		return p.chatDirect(ctx, messages, tools, p.fallbackModel, options, p.fallbackKey, p.fallbackBase)
	}

	return nil, err
}

// isFailoverEligible determines if an error should trigger fallback to the backup model.
// We failover on server errors, timeouts, and overload, but NOT on client errors (400, 401, 403).
func (p *HTTPProvider) isFailoverEligible(err error) bool {
	errStr := err.Error()
	// Failover-worthy errors
	if strings.Contains(errStr, "overloaded") || strings.Contains(errStr, "engine_overloaded") {
		return true
	}
	if strings.Contains(errStr, "timeout") || strings.Contains(errStr, "deadline") {
		return true
	}
	if strings.Contains(errStr, "500") || strings.Contains(errStr, "502") || strings.Contains(errStr, "503") || strings.Contains(errStr, "529") {
		return true
	}
	if strings.Contains(errStr, "failed after") {
		return true // All retries exhausted
	}
	// Do NOT failover on client errors (auth, bad request, content policy)
	if strings.Contains(errStr, "401") || strings.Contains(errStr, "403") {
		return false
	}
	return false
}

// chatDirect performs the actual HTTP call to a specific endpoint.
// This is used by both primary and fallback paths.
func (p *HTTPProvider) chatDirect(ctx context.Context, messages []Message, tools []ToolDefinition, model string, options map[string]interface{}, apiKey, apiBase string) (*LLMResponse, error) {
	// Validate and clean message chain before sending
	cleanMessages := p.validateMessages(messages)

	requestBody := map[string]interface{}{
		"model":    model,
		"messages": cleanMessages,
	}

	if len(tools) > 0 {
		requestBody["tools"] = tools
		requestBody["tool_choice"] = "auto"
	}

	if maxTokens, ok := options["max_tokens"].(int); ok {
		lowerModel := strings.ToLower(model)
		if strings.Contains(lowerModel, "glm") || strings.Contains(lowerModel, "o1") {
			requestBody["max_completion_tokens"] = maxTokens
		} else {
			requestBody["max_tokens"] = maxTokens
		}
	}

	if temperature, ok := options["temperature"].(float64); ok {
		requestBody["temperature"] = temperature
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Retry with exponential backoff for transient errors
	maxRetries := 3
	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(1<<uint(attempt)) * time.Second // 2s, 4s, 8s
			logger.WarnCF("provider", fmt.Sprintf("Retrying API call (attempt %d/%d) after %v", attempt, maxRetries, backoff),
				map[string]interface{}{"attempt": attempt})
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		req, err := http.NewRequestWithContext(ctx, "POST", apiBase+"/chat/completions", bytes.NewReader(jsonData))
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		if apiKey != "" {
			req.Header.Set("Authorization", "Bearer "+apiKey)
		}

		resp, err := p.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to send request: %w", err)
			continue // Network error, retry
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = fmt.Errorf("failed to read response: %w", err)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			return p.parseResponse(body)
		}

		// Check if retryable
		if p.isRetryableStatus(resp.StatusCode, body) && attempt < maxRetries {
			lastErr = fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
			logger.WarnCF("provider", "Transient API error, will retry",
				map[string]interface{}{"status": resp.StatusCode, "body": string(body)})
			continue
		}

		// Non-retryable or final attempt
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	return nil, fmt.Errorf("API call failed after %d retries: %w", maxRetries, lastErr)
}

// isRetryableStatus checks if an API error is transient and worth retrying.
func (p *HTTPProvider) isRetryableStatus(statusCode int, body []byte) bool {
	// HTTP-level retryable statuses
	switch statusCode {
	case 429, 500, 502, 503, 529:
		return true
	}
	// Check for engine_overloaded in response body (Moonshot returns this as 200-level sometimes)
	bodyStr := string(body)
	if strings.Contains(bodyStr, "engine_overloaded") || strings.Contains(bodyStr, "overloaded") {
		return true
	}
	return false
}

func (p *HTTPProvider) parseResponse(body []byte) (*LLMResponse, error) {
	var apiResponse struct {
		Choices []struct {
			Message struct {
				Content   string `json:"content"`
				ToolCalls []struct {
					ID       string `json:"id"`
					Type     string `json:"type"`
					Function *struct {
						Name      string `json:"name"`
						Arguments string `json:"arguments"`
					} `json:"function"`
				} `json:"tool_calls"`
			} `json:"message"`
			FinishReason string `json:"finish_reason"`
		} `json:"choices"`
		Usage *UsageInfo `json:"usage"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(apiResponse.Choices) == 0 {
		return &LLMResponse{
			Content:      "",
			FinishReason: "stop",
		}, nil
	}

	choice := apiResponse.Choices[0]

	toolCalls := make([]ToolCall, 0, len(choice.Message.ToolCalls))
	for _, tc := range choice.Message.ToolCalls {
		arguments := make(map[string]interface{})
		name := ""

		// Handle OpenAI format with nested function object
		if tc.Type == "function" && tc.Function != nil {
			name = tc.Function.Name
			if tc.Function.Arguments != "" {
				if err := json.Unmarshal([]byte(tc.Function.Arguments), &arguments); err != nil {
					arguments["raw"] = tc.Function.Arguments
				}
			}
		} else if tc.Function != nil {
			// Legacy format without type field
			name = tc.Function.Name
			if tc.Function.Arguments != "" {
				if err := json.Unmarshal([]byte(tc.Function.Arguments), &arguments); err != nil {
					arguments["raw"] = tc.Function.Arguments
				}
			}
		}

		newTC := ToolCall{
			ID:        tc.ID,
			Type:      tc.Type,
			Name:      name,
			Arguments: arguments,
		}
		if tc.Function != nil {
			newTC.Function = &FunctionCall{
				Name:      tc.Function.Name,
				Arguments: tc.Function.Arguments,
			}
		}
		toolCalls = append(toolCalls, newTC)
	}

	return &LLMResponse{
		Content:      choice.Message.Content,
		ToolCalls:    toolCalls,
		FinishReason: choice.FinishReason,
		Usage:        apiResponse.Usage,
	}, nil
}

// validateMessages ensures that every 'tool' message has a matching 'assistant' message with that ID.
func (p *HTTPProvider) validateMessages(messages []Message) []Message {
	validIDs := make(map[string]bool)
	validated := make([]Message, 0, len(messages))

	for _, msg := range messages {
		if msg.Role == "assistant" && len(msg.ToolCalls) > 0 {
			for _, tc := range msg.ToolCalls {
				validIDs[tc.ID] = true
			}
			validated = append(validated, msg)
		} else if msg.Role == "tool" {
			if validIDs[msg.ToolCallID] {
				validated = append(validated, msg)
			} else {
				// Skip orphaned tool messages that would cause API errors
				logger.WarnCF("provider", "Stripping orphaned tool message from history", map[string]interface{}{
					"tool_call_id": msg.ToolCallID,
				})
			}
		} else {
			validated = append(validated, msg)
		}
	}
	return validated
}

func (p *HTTPProvider) GetDefaultModel() string {
	return ""
}

func CreateProvider(cfg *config.Config) (LLMProvider, error) {
	model := cfg.Agents.Defaults.Model

	var apiKey, apiBase string

	lowerModel := strings.ToLower(model)

	switch {
	case strings.HasPrefix(model, "openrouter/") || strings.HasPrefix(model, "anthropic/") || strings.HasPrefix(model, "openai/") || strings.HasPrefix(model, "meta-llama/") || strings.HasPrefix(model, "deepseek/") || strings.HasPrefix(model, "google/"):
		apiKey = cfg.Providers.OpenRouter.APIKey
		if cfg.Providers.OpenRouter.APIBase != "" {
			apiBase = cfg.Providers.OpenRouter.APIBase
		} else {
			apiBase = "https://openrouter.ai/api/v1"
		}

	case (strings.Contains(lowerModel, "claude") || strings.HasPrefix(model, "anthropic/")) && cfg.Providers.Anthropic.APIKey != "":
		apiKey = cfg.Providers.Anthropic.APIKey
		apiBase = cfg.Providers.Anthropic.APIBase
		if apiBase == "" {
			apiBase = "https://api.anthropic.com/v1"
		}

	case (strings.Contains(lowerModel, "gpt") || strings.HasPrefix(model, "openai/")) && cfg.Providers.OpenAI.APIKey != "":
		apiKey = cfg.Providers.OpenAI.APIKey
		apiBase = cfg.Providers.OpenAI.APIBase
		if apiBase == "" {
			apiBase = "https://api.openai.com/v1"
		}

	case (strings.Contains(lowerModel, "gemini") || strings.HasPrefix(model, "google/")) && cfg.Providers.Gemini.APIKey != "":
		apiKey = cfg.Providers.Gemini.APIKey
		apiBase = cfg.Providers.Gemini.APIBase
		if apiBase == "" {
			apiBase = "https://generativelanguage.googleapis.com/v1beta/openai"
		}

	case (strings.Contains(lowerModel, "glm") || strings.Contains(lowerModel, "zhipu") || strings.Contains(lowerModel, "zai")) && cfg.Providers.Zhipu.APIKey != "":
		apiKey = cfg.Providers.Zhipu.APIKey
		apiBase = cfg.Providers.Zhipu.APIBase
		if apiBase == "" {
			apiBase = "https://open.bigmodel.cn/api/paas/v4"
		}

	case (strings.Contains(lowerModel, "groq") || strings.HasPrefix(model, "groq/")) && cfg.Providers.Groq.APIKey != "":
		apiKey = cfg.Providers.Groq.APIKey
		apiBase = cfg.Providers.Groq.APIBase
		if apiBase == "" {
			apiBase = "https://api.groq.com/openai/v1"
		}

	case (strings.Contains(lowerModel, "moonshot") || strings.HasPrefix(model, "moonshot/")) && cfg.Providers.Moonshot.APIKey != "":
		apiKey = cfg.Providers.Moonshot.APIKey
		apiBase = cfg.Providers.Moonshot.APIBase
		if apiBase == "" {
			apiBase = "https://api.moonshot.ai/v1"
		}

	case cfg.Providers.VLLM.APIBase != "":
		apiKey = cfg.Providers.VLLM.APIKey
		apiBase = cfg.Providers.VLLM.APIBase

	default:
		if cfg.Providers.OpenRouter.APIKey != "" {
			apiKey = cfg.Providers.OpenRouter.APIKey
			if cfg.Providers.OpenRouter.APIBase != "" {
				apiBase = cfg.Providers.OpenRouter.APIBase
			} else {
				apiBase = "https://openrouter.ai/api/v1"
			}
		} else {
			return nil, fmt.Errorf("no API key configured for model: %s", model)
		}
	}

	if apiKey == "" && !strings.HasPrefix(model, "bedrock/") {
		return nil, fmt.Errorf("no API key configured for provider (model: %s)", model)
	}

	if apiBase == "" {
		return nil, fmt.Errorf("no API base configured for provider (model: %s)", model)
	}

	// Resolve fallback provider if configured
	fallbackModel := cfg.Agents.Defaults.FallbackModel
	if fallbackModel != "" {
		fallbackKey, fallbackBase := resolveFallbackProvider(cfg, fallbackModel)
		if fallbackKey != "" && fallbackBase != "" {
			logger.InfoCF("provider", fmt.Sprintf("Failover configured: %s → %s", model, fallbackModel),
				map[string]interface{}{
					"primary":  model,
					"fallback": fallbackModel,
				})
			return NewHTTPProviderWithFallback(apiKey, apiBase, fallbackKey, fallbackBase, fallbackModel), nil
		}
	}

	return NewHTTPProvider(apiKey, apiBase), nil
}

// resolveFallbackProvider resolves the API key and base URL for the fallback model.
func resolveFallbackProvider(cfg *config.Config, fallbackModel string) (string, string) {
	lower := strings.ToLower(fallbackModel)

	switch {
	case strings.Contains(lower, "gemini"):
		key := cfg.Providers.Gemini.APIKey
		base := cfg.Providers.Gemini.APIBase
		if base == "" {
			base = "https://generativelanguage.googleapis.com/v1beta/openai"
		}
		return key, base

	case strings.Contains(lower, "moonshot"):
		key := cfg.Providers.Moonshot.APIKey
		base := cfg.Providers.Moonshot.APIBase
		if base == "" {
			base = "https://api.moonshot.ai/v1"
		}
		return key, base

	case strings.Contains(lower, "gpt"):
		key := cfg.Providers.OpenAI.APIKey
		base := cfg.Providers.OpenAI.APIBase
		if base == "" {
			base = "https://api.openai.com/v1"
		}
		return key, base

	case strings.Contains(lower, "claude"):
		key := cfg.Providers.Anthropic.APIKey
		base := cfg.Providers.Anthropic.APIBase
		if base == "" {
			base = "https://api.anthropic.com/v1"
		}
		return key, base
	}

	return "", ""
}
