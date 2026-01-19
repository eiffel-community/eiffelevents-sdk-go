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

package validator

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBundledSchemaLocator(t *testing.T) {
	testcases := []struct {
		name          string
		eventType     string
		version       string
		expectSuccess bool
	}{
		{
			name:          "Happy path",
			eventType:     "EiffelCompositionDefinedEvent",
			version:       "3.3.0",
			expectSuccess: true,
		},
		{
			name:          "Asking for non-existent event results in a nil return",
			eventType:     "EiffelBogusEvent",
			version:       "1.0.0",
			expectSuccess: false,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			loc := NewBundledSchemaLocator()
			body, err := loc.GetSchema(t.Context(), tc.eventType, tc.version, "")
			if tc.expectSuccess {
				require.NoError(t, err)
				require.NotNil(t, body)
				b, err := io.ReadAll(body)
				require.NoError(t, err)
				assert.True(t, strings.HasPrefix(string(b), "{"))
			} else {
				assert.Nil(t, body)
				assert.NoError(t, err)
			}
		})
	}
}
