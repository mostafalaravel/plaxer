package headers

import (
    "fmt"
    "strings"
)

// ParseHeaders takes a slice of header strings (e.g., "Authorization: Bearer token")
// and returns a map of headers. It returns an error if any header is invalid.
func ParseHeaders(headerStrings []string) (map[string]string, error) {
    headers := make(map[string]string)

    for _, h := range headerStrings {
        // Split the header string into key and value.
        parts := strings.SplitN(h, ":", 2)
        if len(parts) != 2 {
            return nil, fmt.Errorf("invalid header format: %s", h)
        }

        // Trim spaces from the key and value.
        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])

        // Ensure the key and value are not empty.
        if key == "" || value == "" {
            return nil, fmt.Errorf("header key or value cannot be empty: %s", h)
        }

        // Add the header to the map.
        headers[key] = value
    }

    return headers, nil
}