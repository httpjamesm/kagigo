package kagi

import (
	"os"
	"testing"

	"github.com/httpjamesm/kagigo/constants"
)

func TestFastGPTCompletion(t *testing.T) {
	apiToken := os.Getenv("KAGI_API_TOKEN")

	// Create a new test client
	client := NewClient(&ClientConfig{
		APIKey:     apiToken,
		APIVersion: constants.CurrentApiVersion,
	})

	// Define test input parameters
	params := FastGPTCompletionParams{
		Query:     "test query",
		WebSearch: false,
		Cache:     true,
	}

	// Call the function being tested
	res, err := client.FastGPTCompletion(params)

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the response fields
	if res.Meta.ID == "" {
		t.Errorf("Expected non-empty ID, got '%s'", res.Meta.ID)
	}

	if res.Meta.Node == "" {
		t.Errorf("Expected non-empty Node, got '%s'", res.Meta.Node)
	}

	if res.Meta.Ms < 0 {
		t.Errorf("Expected positive Ms, got '%d'", res.Meta.Ms)
	}

	if res.Data.Output == "" {
		t.Errorf("Expected non-empty Output, got '%s'", res.Data.Output)
	}

	if res.Data.Tokens < 0 {
		t.Errorf("Expected positive Tokens, got '%d'", res.Data.Tokens)
	}
}
