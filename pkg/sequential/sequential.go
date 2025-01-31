package sequential

import (
    "fmt"
    "log"
    "plaxer/internal/checker"
    "plaxer/pkg/runner"
    "time"
)

// RunSequentially sends HTTP requests sequentially to the specified URL.
// It takes the URL, number of requests, headers, HTTP method, payload, and timeout as input.
// If stats is true, it prints HTTP status code statistics after all requests are completed.
func RunSequentially(url string, requests int, stats bool, headers map[string]string, httpMethod string, payload []byte, timeout time.Duration) {
    statusCounts := make(map[int]int) // Map to store counts of each HTTP status code

    for i := 0; i < requests; i++ {
        // Perform the HTTP request and get the status code and response time.
        statusCode, responseTime, err := checker.CheckHTTPStatus(httpMethod, url, headers, payload, timeout)
        if err != nil {
            log.Printf("Request %d: Failed to fetch URL: %v\n", i+1, err)
            continue
        }

        // Increment the count for the returned status code.
        statusCounts[statusCode]++
        fmt.Printf("Request %d: HTTP Status Code: %d, Response Time: %dms\n", i+1, statusCode, responseTime)
    }

    // Print statistics if enabled.
    if stats {
        runner.PrintStatistics(statusCounts, requests)
    }
}