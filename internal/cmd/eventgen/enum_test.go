// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
