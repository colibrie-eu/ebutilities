package utils

import (
	"errors"
	"os"
	"time"
)

const (
	Backoff = 100
)

var ErrFileEmpty = errors.New("file is empty")

// ReadFile reads the file with the given filename and returns its content.
func ReadFile(filename string, maxRetries int) ([]byte, error) {
	var lastErr error

	for attempt := range maxRetries {
		read, err := os.ReadFile(filename)
		if err == nil {
			// Success - but check if file has content
			if len(read) > 0 {
				return read, nil
			}
			// File exists but empty, treat as not ready yet
			lastErr = ErrFileEmpty
		} else {
			lastErr = err
		}

		// Don't sleep on the last attempt
		if attempt < maxRetries-1 {
			// Exponential backoff: 100ms, 200ms, 400ms, 800ms, etc.
			backoffDuration := time.Duration(Backoff*(1<<attempt)) * time.Millisecond
			time.Sleep(backoffDuration)
		}
	}

	return nil, lastErr
}
