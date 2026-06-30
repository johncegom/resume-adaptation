package adaptation

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/genai"
)

const (
	defaultModel   = "gemini-2.5-flash"
	defaultTimeout = 30 * time.Second
	apiKeyEnvVar   = "GEMINI_API_KEY"
)

// clientFactory creates a genai.Client. Extracted for test injection.
type clientFactory func(ctx context.Context, config *genai.ClientConfig) (*genai.Client, error)

// defaultClientFactory calls the real Gen AI SDK.
func defaultClientFactory(ctx context.Context, config *genai.ClientConfig) (*genai.Client, error) {
	return genai.NewClient(ctx, config)
}

// Client wraps the Google Gen AI SDK client with project-specific
// configuration for model selection and timeout enforcement.
type Client struct {
	inner   *genai.Client
	model   string
	timeout time.Duration
	factory clientFactory
}

// Option configures a Client during construction.
type Option func(*Client)

// WithModel sets the Gemini model name (e.g. "gemini-2.5-pro").
func WithModel(model string) Option {
	return func(c *Client) { c.model = model }
}

// WithTimeout sets the context timeout for API calls.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.timeout = d }
}

// withFactory sets the client factory (for testing only).
func withFactory(f clientFactory) Option {
	return func(c *Client) { c.factory = f }
}

// NewClient creates a new adaptation Client by reading the API key
// from the GEMINI_API_KEY environment variable and initializing the
// official Gen AI SDK client.
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	apiKey := os.Getenv(apiKeyEnvVar)
	if apiKey == "" {
		return nil, fmt.Errorf("environment variable %s is not set", apiKeyEnvVar)
	}

	c := &Client{
		model:   defaultModel,
		timeout: defaultTimeout,
		factory: defaultClientFactory,
	}
	for _, opt := range opts {
		opt(c)
	}

	inner, err := c.factory(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}
	c.inner = inner

	return c, nil
}

// Model returns the configured model name.
func (c *Client) Model() string {
	return c.model
}

// Timeout returns the configured timeout duration.
func (c *Client) Timeout() time.Duration {
	return c.timeout
}
