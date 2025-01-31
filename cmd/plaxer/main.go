package main

import (
    "log"
    "os"
    "plaxer/pkg/headers"
    "plaxer/pkg/logger"
    "plaxer/pkg/sequential"
    "plaxer/pkg/concurrent"
    "time"
    "fmt"

    "github.com/joho/godotenv" // Import the godotenv package
)

// Load environment variables from the .env file.
func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

// GetAppInfo returns the application name, version, and author from the .env file.
func GetAppInfo() map[string]string {
    return map[string]string{
        "AppName":    os.Getenv("APP_NAME"),
        "AppVersion": os.Getenv("APP_VERSION"),
        "AppAuthor":  os.Getenv("APP_AUTHOR"),
        "AppDescription": os.Getenv("APP_DESCRIPTION"),
        "AppLicense": os.Getenv("APP_LICENSE"),
        "AppRepository": os.Getenv("APP_REPOSITORY"),
    }
}

func main() {
    // Initialize the logger.
    logger.Init()

    // Parse command-line arguments.
    args := ParseArgs()

    // If --version is set, print version information and exit.
    if args.Version {
        appInfo := GetAppInfo()
        fmt.Printf("%s v%s\n", appInfo["AppName"], appInfo["AppVersion"])
        fmt.Printf("Author: %s\n", appInfo["AppAuthor"])
        fmt.Printf("Description: %s\n", appInfo["AppDescription"])
        fmt.Printf("License: %s\n", appInfo["AppLicense"])
        fmt.Printf("Repository: %s\n", appInfo["AppRepository"])
        os.Exit(0)
    }

    // Parse headers from the command-line arguments.
    headers, err := headers.ParseHeaders(args.Headers)
    if err != nil {
        log.Fatalf("Error parsing headers: %v", err)
    }

    // Convert the payload to a byte slice.
    payload := []byte(args.Payload)

    // Convert the timeout to a time.Duration.
    timeout := time.Duration(args.Timeout) * time.Second

    // Run requests either concurrently or sequentially based on the --concurrent flag.
    if args.Concurrent {
        concurrent.RunConcurrently(args.Url, args.Requests, args.Stats, headers, args.HttpMethod, payload, timeout)
    } else {
        sequential.RunSequentially(args.Url, args.Requests, args.Stats, headers, args.HttpMethod, payload, timeout)
    }
}