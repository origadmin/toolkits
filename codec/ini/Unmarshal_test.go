package ini

import (
	"testing"
)

// Successfully unmarshals valid INI data into the provided object
func TestUnmarshalValidINIData(t *testing.T) {
	// Arrr, let's see if this INI treasure can be mapped to our struct!
	data := []byte("[section]\nkey=value")
	var result struct {
		Section struct {
			Key string `ini:"key"`
		} `ini:"section"`
	}
	err := Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if result.Section.Key != "value" {
		t.Fatalf("Expected 'value', but got %v", result.Section.Key)
	}
}
