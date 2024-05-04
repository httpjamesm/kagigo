package kagi

import (
	"os"
	"testing"

	"github.com/httpjamesm/kagigo/constants"
)

func TestEnrichmentCompletion(t *testing.T) {
	apiToken := os.Getenv("KAGI_API_TOKEN")

	client := NewClient(&ClientConfig{
		APIKey:     apiToken,
		APIVersion: constants.CurrentApiVersion,
	})

	params := EnrichmentParams{
		Q: "kagi search",
	}

	// test EndpointType exists error
	res, err := client.EnrichmentCompletion("", params)
	if err == nil {
		t.Errorf("Expected endpoint type is required error, got nil")
	}

	// test EndpointType must be web or news error
	res, err = client.EnrichmentCompletion("unexpected", params)
	if err == nil {
		t.Errorf("Expected endpoint type must be web or news error, got nil")
	}

	// test news endpoint does not return error
	res, err = client.EnrichmentCompletion("news", EnrichmentParams{
		Q: "los angeles",
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// test web endpoint does not return error
	res, err = client.EnrichmentCompletion("web", params)
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

	if res.Data == nil {
		t.Errorf("Expected non-empty Data object, got '%v'", res.Data)
	}

	if res.Data[0].URL == "" || res.Data[0].Title == "" || res.Data[0].Snippet == "" {
		t.Errorf("Expected URL, Title, Snippet, got '%v'", res.Data[0])
	}

	// test if query is empty
	params.Q = ""
	res, err = client.EnrichmentCompletion("web", params)
	if err == nil {
		t.Errorf("Expected query is required error, got nil")
	}
}
