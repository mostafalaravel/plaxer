package httpclient

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// TestRequest tests the Request function.
// It ensures that the function works correctly for both GET and POST requests.
func TestRequest(t *testing.T) {
    // Create a mock HTTP server.
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))
    defer server.Close()

    // Test a GET request.
    resp, err := Request(http.MethodGet, server.URL, nil, nil, 10*time.Second)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if resp.StatusCode != http.StatusOK {
        t.Errorf("expected status 200, got %d", resp.StatusCode)
    }

    // Test a POST request with headers and body.
    headers := map[string]string{"Content-Type": "application/json"}
    body := []byte(`{"key": "value"}`)
    resp, err = Request(http.MethodPost, server.URL, headers, body, 10*time.Second)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if resp.StatusCode != http.StatusOK {
        t.Errorf("expected status 200, got %d", resp.StatusCode)
    }
}