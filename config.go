package openaifx

import "net/http"

// Config holds the OpenAI client configuration.
// APIKey is the OpenAI API key (required).
// BaseURL is an optional custom API endpoint URL.
// HTTPClient is an optional custom HTTP client.
type Config struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}
