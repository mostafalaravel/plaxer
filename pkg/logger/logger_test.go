package logger

import (
    "testing"
)

// TestLoggerInit tests the Init function.
// It ensures that the logger initialization does not panic.
func TestLoggerInit(t *testing.T) {
    Init()
}