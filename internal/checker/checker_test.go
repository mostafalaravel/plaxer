package checker

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// TestCheckHTTPStatus tests the CheckHTTPStatus function.
// It ensures that the function works correctly for valid and invalid URLs.
func TestCheckHTTPStatus(t *testing.T) {
    // Create a mock HTTP server.
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))
    defer server.Close()

    // Test a successful request.
    status, duration, err := CheckHTTPStatus(http.MethodGet, server.URL, nil, nil, 10*time.Second)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if status != http.StatusOK {
        t.Errorf("expected status 200, got %d", status)
    }
    if duration < 0 {
        t.Errorf("expected non-negative duration, got %d", duration)
    }

    // Test an invalid URL.
    _, _, err = CheckHTTPStatus(http.MethodGet, "invalid-url", nil, nil, 10*time.Second)
    if err == nil {
        t.Error("expected error for invalid URL, got nil")
    }
}