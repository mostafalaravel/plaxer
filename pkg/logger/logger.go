package logger

import (
    "log"
    "os"
)

// Init initializes the logger with default settings.
// It sets the output to stdout and configures the log format to include date, time, and file information.
func Init() {
    log.SetOutput(os.Stdout)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}