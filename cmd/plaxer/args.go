package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
)

type Arguments struct {
    Url         string   // URL to send requests to
    Requests    int      // Number of requests to send
    Concurrent  bool     // Whether to run requests concurrently
    Stats       bool     // Whether to display HTTP status code statistics
    Headers     []string // Custom headers to include in requests
    HttpMethod  string   // HTTP method to use (e.g., GET, POST)
    Payload     string   // Request payload for POST/PUT requests
    Timeout     int      // Timeout for each request in seconds
    Version     bool     // Whether to display version information
}

type headersFlags []string

func (h *headersFlags) String() string {
    return strings.Join(*h, ", ")
}

func (h *headersFlags) Set(value string) error {
    *h = append(*h, value)
    return nil
}

func ParseArgs() *Arguments {
    // Define command-line flags.
    url := flag.String("url", "", "URL to check (required)")
    requests := flag.Int("req-num", 1, "Number of times to repeat the request (default: 1)")
    concurrent := flag.Bool("concurrent", false, "Run requests concurrently (default: false)")
    stats := flag.Bool("stats", false, "Display HTTP status code stats (default: false)")
    httpMethod := flag.String("http-method", "GET", "HTTP method to use (e.g., GET, POST, PUT, DELETE)")
    payload := flag.String("payload", "", "Request payload for POST/PUT requests (e.g., '{\"name\": \"mostafa\"}')")
    timeout := flag.Int("timeout", 10, "Timeout for each request in seconds (default: 10)")
    version := flag.Bool("version", false, "Display version information") // New flag

    var headers headersFlags
    flag.Var(&headers, "headers", "Custom headers to include in the request (e.g., 'Authorization: Bearer token')")

    // Customize the usage message.
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: plaxer --url <URL> [--req-num <n>] [--concurrent] [--stats] [--http-method <http-method>] [--payload <payload>] [--headers <headers>]\n")
        fmt.Fprintf(os.Stderr, "Options:\n")
        flag.PrintDefaults()
    }

    // Parse the command-line arguments.
    flag.Parse()

    // Show usage if no arguments are provided.
    if len(os.Args) == 1 {
        flag.Usage()
        os.Exit(1)
    }

    // If --version is set, print version information and exit.
    if *version {
        appInfo := GetAppInfo()
        fmt.Printf("%s v%s\n", appInfo["AppName"], appInfo["AppVersion"])
        fmt.Printf("Author: %s\n", appInfo["AppAuthor"])
        fmt.Printf("Description: %s\n", appInfo["AppDescription"])
        fmt.Printf("License: %s\n", appInfo["AppLicense"])
        fmt.Printf("Repository: %s\n", appInfo["AppRepository"])
        os.Exit(0)
    }

    // Validate required flags.
    if *url == "" {
        log.Fatal("Error: --url flag is required")
    }
    if *requests < 1 {
        log.Fatal("Error: --req-num must be at least 1")
    }

    // Validate the HTTP method.
    validHttpMethods := map[string]bool{
        "GET":    true,
        "POST":   true,
        "PUT":    true,
        "DELETE": true,
        "PATCH":  true,
    }
    if !validHttpMethods[strings.ToUpper(*httpMethod)] {
        log.Fatalf("Error: Invalid HTTP method '%s'. Supported methods are GET, POST, PUT, DELETE, PATCH", *httpMethod)
    }

    // Return the parsed arguments.
    return &Arguments{
        Url:        *url,
        Requests:   *requests,
        Concurrent: *concurrent,
        Stats:      *stats,
        Headers:    headers,
        HttpMethod: strings.ToUpper(*httpMethod),
        Payload:    *payload,
        Timeout:    *timeout,
        Version:    *version,
    }
}