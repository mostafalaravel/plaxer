package concurrent

import (
    "fmt"
    "plaxer/internal/checker"
    "plaxer/pkg/runner"
    "sync"
    "time"
)

// RunConcurrently sends HTTP requests concurrently to the specified URL.
// It takes the URL, number of requests, headers, HTTP method, payload, and timeout as input.
// If stats is true, it prints HTTP status code statistics after all requests are completed.
func RunConcurrently(url string, requests int, stats bool, headers map[string]string, httpMethod string, payload []byte, timeout time.Duration) {
    var wg sync.WaitGroup
    statusCounts := make(map[int]int)      // Map to store counts of each HTTP status code
    statusCountsMutex := sync.Mutex{}     // Mutex to protect concurrent access to statusCounts

    for i := 0; i < requests; i++ {
        wg.Add(1)
        go func(requestNumber int) {
            defer wg.Done()

            // Perform the HTTP request and get the status code and response time.
            statusCode, responseTime, err := checker.CheckHTTPStatus(httpMethod, url, headers, payload, timeout)
            if err != nil {
                // Truncate the error message if it's too long.
                errMsg := err.Error()
                if len(errMsg) > 20 {
                    errMsg = errMsg[:20] + "..."
                }
                fmt.Printf("Request %d: Failed to fetch URL: %s\n", requestNumber, errMsg)
                return
            }

            // Safely increment the count for the returned status code.
            statusCountsMutex.Lock()
            statusCounts[statusCode]++
            statusCountsMutex.Unlock()

            fmt.Printf("Request %d: HTTP Status Code: %d, Response Time: %dms\n", requestNumber, statusCode, responseTime)
        }(i + 1)
    }

    // Wait for all goroutines to finish.
    wg.Wait()

    // Print statistics if enabled.
    if stats {
        runner.PrintStatistics(statusCounts, requests)
    }
}