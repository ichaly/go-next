package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartWithAny(t *testing.T) {
	tests := []struct {
		s            string
		prefixes     []string
		expected     string
		expectedBool bool
	}{
		{"hello world", []string{"he"}, "he", true},
		{"hello world", []string{"wo"}, "", false},
		{"hello world", []string{"h", "e"}, "h", true},
		{"hello world", []string{"hello", "world"}, "hello", true},
	}

	for _, tt := range tests {
		actual, actualBool := StartWithAny(tt.s, tt.prefixes...)
		assert.Equal(t, tt.expected, actual, "Expected prefix not found")
		assert.Equal(t, tt.expectedBool, actualBool, "Expected boolean value not found")
	}
}
