package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"ApiKey 1234dfghn"},
	}
	apiKey, err := GetAPIKey(header)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if apiKey != "1234dfghn" {
		t.Errorf("Expected API key '1234dfghn', got %s", apiKey)
	}

	emptyHeaders := http.Header{}
	_, err = GetAPIKey(emptyHeaders)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
	}

	malformedHeaders := http.Header{"Authorization": []string{"Bearer invalid-token"}}
	_, err = GetAPIKey(malformedHeaders)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("Expected 'malformed authorization header' error, got %v", err)
	}
}
