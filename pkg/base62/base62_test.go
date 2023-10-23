package base62

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase62_Encode(t *testing.T) {
	testCases := []struct {
		n        int64
		expected string
	}{
		{0, "0"},
		{10, "a"},
		{36, "A"},
		{61, "Z"},
		{62, "01"},
		{9223372036854775807, "7M85y0N8lZa"},
	}
	base62 := NewConverter()
	for _, tc := range testCases {
		result := base62.Encode(tc.n)
		assert.Equal(t, tc.expected, result, fmt.Sprintf("Expected %s but got %d", tc.expected, tc.n))
	}
}

func TestBase62_Decode(t *testing.T) {
	testCases := []struct {
		s        string
		expected int64
	}{
		{"0", 0},
		{"a", 10},
		{"A", 36},
		{"Z", 61},
		{"01", 62},
		{"7M85y0N8lZa", 9223372036854775807},
	}
	base62 := NewConverter()
	for _, tc := range testCases {
		result := base62.Decode(tc.s)
		assert.Equal(t, tc.expected, result, fmt.Sprintf("Expected %d but got %s", tc.expected, tc.s))
	}
}
