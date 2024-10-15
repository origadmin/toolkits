package replacer

import (
	"strings"

	"github.com/goexts/ggb/settings"
)

// Matcher interface defines methods for matching and replacing strings.
type Matcher interface {
	Match(content string) (string, bool)
	Replace(content string) string
}
type matcher struct {
	offset      int
	sta         string
	end         string
	fold        bool
	replacement map[string]string
}

func (m matcher) Match(content string) (string, bool) {
	cursor := 0
	for {
		// Find the next occurrence of `${`
		sta := strings.Index(content[cursor:], m.sta)
		if sta == -1 {
			// No more occurrences, write the remaining content and break
			break
		}

		// Find the closing `}`
		end := strings.Index(content[cursor+sta:], m.end)
		if end == -1 {
			// No closing brace found, write the remaining content and break
			break
		}
		// Extract the variable name
		varName := content[cursor+sta+m.offset : cursor+sta+end]
		// Check for replacement in the map (case-insensitive)
		for key, value := range m.replacement {
			if defaultMatchFunc(varName, key, m.fold) {
				return value, true
			}
		}
		// Update the cursor position for the next iteration
		cursor = cursor + sta + end + 1
	}

	return "", false
}

func (m matcher) Replace(content string) string {
	cursor := 0
	var sb strings.Builder

	for {
		// Find the next occurrence of `${`
		sta := strings.Index(content[cursor:], m.sta)
		if sta == -1 {
			sb.WriteString(content[cursor:])
			// No more occurrences, write the remaining content and break
			break
		}

		// Write the content before the found pattern
		sb.WriteString(content[cursor : cursor+sta])

		// Find the closing `}`
		end := strings.Index(content[cursor+sta:], m.end)
		if end == -1 {
			sb.WriteString(content[cursor+sta:])
			// No closing brace found, write the remaining content and break
			break
		}
		// Extract the variable name
		varName := content[cursor+sta+m.offset : cursor+sta+end]
		// Check for replacement in the map (case-insensitive)
		found := false
		for key, value := range m.replacement {
			if value != "" && defaultMatchFunc(varName, key, m.fold) {
				sb.WriteString(value)
				found = true
				break
			}
		}

		if !found {
			sb.WriteString(m.sta + varName + m.end)
		}
		// Update the cursor position for the next iteration
		cursor = cursor + sta + end + 1
	}

	return sb.String()
}

// NewMatch creates a new matcher with the provided replacements.
func NewMatch(replacements map[string]string, ss ...MatchSetting) Matcher {
	m := settings.Apply(&matcher{
		replacement: replacements,
		sta:         DefaultMatchStartKeyword,
		end:         DefaultMatchEndKeyword,
	}, ss)

	m.offset = len(m.sta)
	return m
}
