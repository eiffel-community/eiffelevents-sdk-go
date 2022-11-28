package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToEnu(t *testing.T) {
	testcases := []struct {
		input          string
		expectedOutput string
	}{
		{"PLAIN", "Plain"},
		{"MULTIPLE_WORDS", "MultipleWords"},
		{"ABBREV-123", "ABBREV_123"},
		{"ABBREV-123/456", "ABBREV_123_456"},
		{"SHA256", "SHA256"},
	}
	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expectedOutput, stringToEnum(tc.input))
		})
	}
}
