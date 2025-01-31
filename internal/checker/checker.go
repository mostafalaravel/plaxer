package checker

import (
    "fmt"
    "plaxer/internal/httpclient"
    "time"
)

// CheckHTTPStatus performs an HTTP request and returns the status code and response time.
// It takes the HTTP method, URL, headers, body, and timeout as input.
func CheckHTTPStatus(method, url string, headers map[string]string, body []byte, timeout time.Duration) (int, int64, error) {
    start := time.Now()

    // Perform the HTTP request.
    resp, err := httpclient.Request(method, url, headers, body, timeout)
    if err != nil {
        return 0, 0, fmt.Errorf("failed to perform HTTP request: %v", err)
    }
    defer resp.Body.Close()

    // Calculate the response time in microseconds.
    duration := time.Since(start).Microseconds()
    return resp.StatusCode, duration, nil
}