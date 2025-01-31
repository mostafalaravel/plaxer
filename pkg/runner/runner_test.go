package runner

import (
    "testing"
)

// TestPrintStatistics tests the PrintStatistics function.
// It ensures that the function does not panic when given valid input.
func TestPrintStatistics(t *testing.T) {
    statusCounts := map[int]int{
        200: 8, // 200 OK
        404: 2, // 404 Not Found
    }
    totalRequests := 10

    // Call PrintStatistics and ensure it doesn't panic.
    PrintStatistics(statusCounts, totalRequests)
}