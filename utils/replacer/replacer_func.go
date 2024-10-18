package replacer

import (
	"encoding/json"
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

// FileMatchReplacer replaces occurrences of ${name} with values from the map
func FileMatchReplacer(path string, m Matcher) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	replaced := m.Replace(string(content))
	return []byte(replaced), nil
}

// ObjectReplacer replaces occurrences of ${name} with values from the map
func ObjectReplacer(v any, replacements map[string]string) error {
	if v == nil {
		return nil
	}
	marshal, err := json.Marshal(v)
	if err != nil {
		return err
	}
	bytes := _globalReplacer.Replace(marshal, replacements)
	return json.Unmarshal(bytes, v)
}

// ObjectMatchReplacer replaces occurrences of ${name} with values from the map
func ObjectMatchReplacer(v any, m Matcher) error {
	if v == nil {
		return nil
	}
	marshal, err := json.Marshal(v)
	if err != nil {
		return err
	}
	replaced := m.Replace(string(marshal))
	return json.Unmarshal([]byte(replaced), v)
}
