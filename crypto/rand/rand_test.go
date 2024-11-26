/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package rand

import (
	"testing"
)

func TestKindString(t *testing.T) {
	testCases := []struct {
		kind     Kind
		expected string
	}{
		{Kind(0), "<nil>"},
		{KindDigit, "Digit"},
		{KindLowerCase, "LowerCase"},
		{KindUpperCase, "UpperCase"},
		{KindSymbol, "Symbol"},
		{KindDigit | KindLowerCase, "Digit|LowerCase"},
		{KindDigit | KindLowerCase | KindUpperCase, "Digit|LowerCase|UpperCase"},
		{KindAll, "All"},
		{KindAll | KindDigit, "All"},
		{KindAll | KindDigit | KindLowerCase, "All"},
		{KindAll | KindDigit | KindLowerCase | KindUpperCase, "All"},
		{KindAll | KindDigit | KindLowerCase | KindUpperCase | KindSymbol, "All"},
		{Kind(123), "Custom"},
	}

	for _, tc := range testCases {
		result := tc.kind.String()
		if result != tc.expected {
			t.Errorf("Expected %s but got %s for Kind %d", tc.expected, result, tc.kind)
		}
	}
}
