package runner

import "fmt"

// Define color codes for different HTTP status code ranges.
const (
    Reset   = "\033[0m"  // Reset color to default
    Red     = "\033[31m" // Red color for 4xx status codes
    Yellow  = "\033[33m" // Yellow color for 3xx status codes
    Green   = "\033[32m" // Green color for 2xx status codes
    Magenta = "\033[35m" // Magenta color for 5xx status codes
)

// PrintStatistics prints HTTP status code statistics with color-coded output.
// It takes a map of status codes and their counts, as well as the total number of requests.
func PrintStatistics(statusCounts map[int]int, totalRequests int) {
    fmt.Println("\nHTTP Status Code Statistics:")
    for code, count := range statusCounts {
        percentage := (float64(count) / float64(totalRequests)) * 100

        // Determine the color based on the HTTP status code range.
        var color string
        switch {
        case code >= 200 && code < 300:
            color = Green // 2xx: Success
        case code >= 300 && code < 400:
            color = Yellow // 3xx: Redirection
        case code >= 400 && code < 500:
            color = Red // 4xx: Client errors
        case code >= 500 && code < 600:
            color = Magenta // 5xx: Server errors
        default:
            color = Reset // Unknown status codes
        }

        // Print the status code and its percentage with the appropriate color.
        fmt.Printf("%s[%d] %.2f%%%s\n", color, code, percentage, Reset)
    }
}