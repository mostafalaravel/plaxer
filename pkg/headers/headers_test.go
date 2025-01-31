package headers

import (
    "testing"
)

// TestParseHeaders tests the ParseHeaders function.
// It ensures that valid headers are parsed correctly and invalid headers return errors.
func TestParseHeaders(t *testing.T) {
    // Test valid headers.
    headerStrings := []string{"Authorization: Bearer token", "Content-Type: application/json"}
    headers, err := ParseHeaders(headerStrings)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    // Check if the headers were parsed correctly.
    if headers["Authorization"] != "Bearer token" {
        t.Errorf("expected Authorization header 'Bearer token', got '%s'", headers["Authorization"])
    }
    if headers["Content-Type"] != "application/json" {
        t.Errorf("expected Content-Type header 'application/json', got '%s'", headers["Content-Type"])
    }

    // Test invalid header format.
    _, err = ParseHeaders([]string{"InvalidHeader"})
    if err == nil {
        t.Error("expected error for invalid header format, got nil")
    }

    // Test empty header key.
    _, err = ParseHeaders([]string{": value"})
    if err == nil {
        t.Error("expected error for empty header key, got nil")
    }

    // Test empty header value.
    _, err = ParseHeaders([]string{"key:"})
    if err == nil {
        t.Error("expected error for empty header value, got nil")
    }
}