package replacer

import (
	"os"
)

// FileReplacer opens a file, replaces occurrences of ${name} with values from the map, and returns the result as []byte.
func FileReplacer(path string, replacements map[string]string) ([]byte, error) {
	// Read the file content
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	bytes := _globalReplacer.Replace(content, replacements)
	return bytes, nil
}
