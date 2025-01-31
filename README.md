Plaxer
======

Plaxer is a lightweight, high-performance HTTP benchmarking tool written in Go. It allows you to send sequential or concurrent HTTP requests to a target URL, collect statistics, and analyze the results. Plaxer is designed to be simple, fast, and extensible, making it a great tool for developers and testers who need to evaluate the performance of web services.

Features
--------

*   **Sequential and Concurrent Requests**: Send requests one at a time or concurrently to simulate real-world traffic.
    
*   **Customizable Headers**: Add custom HTTP headers to your requests.
    
*   **HTTP Method Support**: Supports common HTTP methods like `GET`, `POST`, `PUT`, `DELETE`, and `PATCH`.
    
*   **Request Payload**: Send payloads with `POST` or `PUT` requests.
    
*   **Statistics**: Get detailed statistics about HTTP status codes and response times.
    
*   **Color-Coded Output**: Visualize HTTP status codes with color-coded output for better readability.
    
*   **Extensible**: Easily extendable to add new features or customize behavior.
    

Installation
------------

To use Plaxer, you need to have Go installed on your machine. Follow these steps to install and run Plaxer:

1.  **Clone the repository**:
    
    git clone https://github.com/mostafalaravel/plaxer.git
    cd plaxer
    
2.  **Build the project**:
    
    go build \-o plaxer cmd/plaxer/main.go
    
3.  **Run Plaxer**:
    
    ./plaxer \--url <URL\> \[options\]
    

Usage
-----

### Basic Example

Send 10 sequential `GET` requests to a URL:

./plaxer \--url https://example.com --req-num 10

### Concurrent Requests

Send 100 concurrent `GET` requests to a URL:

./plaxer \--url https://example.com --req-num 100 \--concurrent

### Custom Headers and Payload

Send a `POST` request with custom headers and a JSON payload:

./plaxer \--url https://example.com/api --http-method POST \--headers "Content-Type: application/json" \--headers "Authorization: Bearer token" \--payload '{"key": "value"}'

### Enable Statistics

Display HTTP status code statistics after the requests:

./plaxer \--url https://example.com --req-num 50 \--stats

### Full Options

| Flag          | Description                                        | Default Value |
|--------------|----------------------------------------------------|--------------|
| `--url`      | Target URL to send requests to (required).         | -            |
| `--req-num`  | Number of requests to send.                        | `1`          |
| `--concurrent` | Run requests concurrently.                       | `false`      |
| `--stats`    | Display HTTP status code statistics.               | `false`      |
| `--http-method` | HTTP method to use (`GET`, `POST`, `PUT`, `DELETE`, `PATCH`). | `GET` |
| `--payload`  | Request payload for `POST` or `PUT` requests.      | -            |
| `--headers`  | Custom headers to include in the request (can be used multiple times). | - |
| `--version`  | Display the current version | - |


Project Structure
-----------------

```
plaxer/
├── cmd/
│   └── plaxer/
│       ├── main.go          # Entry point for the CLI
│       └── args.go          # Command-line argument parsing
├── pkg/
│   ├── concurrent/          # Concurrent request handling
│   ├── headers/             # Header parsing and validation
│   ├── logger/              # Logging utilities
│   ├── runner/              # Statistics and output formatting
│   └── sequential/          # Sequential request handling
├── internal/
│   ├── checker/             # HTTP status checking
│   └── httpclient/          # HTTP client implementation
├── go.mod                   # Go module file
└── README.md                # Project documentation
```

Contributing
------------

We welcome contributions to Plaxer! If you'd like to contribute, please follow these steps:

1.  Fork the repository.
    
2.  Create a new branch for your feature or bugfix.
    
3.  Make your changes and ensure all tests pass.
    
4.  Submit a pull request with a detailed description of your changes.
    

### Reporting Issues

If you encounter any issues or have suggestions for improvements, please open an issue on the [GitHub Issues](https://github.com/mostafalaravel/plaxer/issues) page.

License
-------

Plaxer is open-source software licensed under the [MIT License](LICENSE).