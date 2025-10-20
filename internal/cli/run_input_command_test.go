package cli

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		input string 
		expected string
	} {
		{input: "01-10-2020", expected: "2020-10-01"},
		{input: "January 2, 2002", expected: "2002-01-02"},
	}

	for _, test := range tests {
		result, _ := parseDate(test.input)
		if result != test.expected {
			t.Errorf("parseDate(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
