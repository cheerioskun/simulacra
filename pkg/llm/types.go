package llm

import "context"

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest represents the request parameters for chat completion
type ChatRequest struct {
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	Temperature float32   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
}

// ChatResponse represents the response from a chat completion
type ChatResponse struct {
	Content string
	Usage   TokenUsage
}

// TokenUsage tracks token usage for the request
type TokenUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Provider defines the interface that all LLM providers must implement
type Provider interface {
	ChatCompletion(ctx context.Context, req ChatRequest) (*ChatResponse, error)
	Name() string
}
