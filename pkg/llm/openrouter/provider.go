package openrouter

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"

	"simulacra/pkg/llm"
)

const defaultAPIEndpoint = "https://openrouter.ai/api/v1"

type Provider struct {
	client *openai.Client
}

func New(apiKey string) *Provider {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = defaultAPIEndpoint

	return &Provider{
		client: openai.NewClientWithConfig(config),
	}
}

func (p *Provider) Name() string {
	return "openrouter"
}

func (p *Provider) ChatCompletion(ctx context.Context, req llm.ChatRequest) (*llm.ChatResponse, error) {
	messages := make([]openai.ChatCompletionMessage, len(req.Messages))
	for i, msg := range req.Messages {
		messages[i] = openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	resp, err := p.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       req.Model,
			Messages:    messages,
			Temperature: float32(req.Temperature),
			MaxTokens:   req.MaxTokens,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("openrouter chat completion failed: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no choices in response")
	}

	return &llm.ChatResponse{
		Content: resp.Choices[0].Message.Content,
		Usage: llm.TokenUsage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}, nil
}
