package openaifx_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/go-core-fx/openaifx"
)

func TestNew_EmptyAPIKey(t *testing.T) {
	_, err := openaifx.New(openaifx.Config{})
	if !errors.Is(err, openaifx.ErrInvalidConfig) {
		t.Fatalf("expected ErrInvalidConfig, got %v", err)
	}
}

func TestNew_ValidAPIKey(t *testing.T) {
	client, err := openaifx.New(openaifx.Config{APIKey: "sk-test"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestNew_CustomBaseURL(t *testing.T) {
	client, err := openaifx.New(openaifx.Config{
		APIKey:  "sk-test",
		BaseURL: "http://localhost:9999/v1/",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestNew_CustomHTTPClient(t *testing.T) {
	custom := &http.Client{Timeout: 5 * time.Second}
	client, err := openaifx.New(openaifx.Config{
		APIKey:     "sk-test",
		HTTPClient: custom,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}
