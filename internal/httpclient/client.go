package httpclient

import (
    "bytes"
    "net/http"
    "time"
)

// Request performs an HTTP request with the specified method, headers, and body.
// It returns the HTTP response and an error if the request fails.
func Request(method, url string, headers map[string]string, body []byte, timeout time.Duration) (*http.Response, error) {
    // Create an HTTP client with the specified timeout.
    client := &http.Client{
        Timeout: timeout,
    }

    // Create a new HTTP request with the specified method, URL, and body.
    req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }

    // Add headers to the request.
    for key, value := range headers {
        req.Header.Add(key, value)
    }

    // Perform the request and return the response.
    return client.Do(req)
}