package utils

import (
	"os"
)

// ReadFile reads the file with the given filename and returns its content.
func ReadFile(filename string) ([]byte, error) {
	read, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return read, nil
}
