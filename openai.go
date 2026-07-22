package openaifx

import (
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

// New creates a new OpenAI client from the provided configuration.
func New(config Config) (*openai.Client, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("%w: API key is empty", ErrInvalidConfig)
	}

	var opts []option.RequestOption

	opts = append(opts, option.WithAPIKey(config.APIKey))

	if config.BaseURL != "" {
		opts = append(opts, option.WithBaseURL(config.BaseURL))
	}

	if config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(config.HTTPClient))
	}

	client := openai.NewClient(opts...)

	return &client, nil
}
