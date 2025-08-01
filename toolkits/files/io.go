package files

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Scanner scans the given reader line by line and calls the given function
func Scanner(r io.Reader, fn func(string) string) *bytes.Buffer {
	buf := new(bytes.Buffer)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		buf.WriteString(fn(line))
		buf.WriteString("\n")
	}
	return buf
}

// IsExists checks if the given file exists
func IsExists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// IsDir Checks if the given path is a directory
func IsDir(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// WriteTo writes the given data to the given file
func WriteTo(name string, data []byte) error {
	return os.WriteFile(name, data, 0644)
}
